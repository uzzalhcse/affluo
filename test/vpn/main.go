package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// DeviceLimits defines restrictions for device usage
type DeviceLimits struct {
	MaxDevicesPerIP      int           // Maximum number of distinct devices per IP
	MaxIPsPerDevice      int           // Maximum number of IPs a device can use
	DeviceTimeout        time.Duration // How long before device session expires
	SuspiciousThreshold  float64       // Risk score threshold for suspicious activity
	BlockThreshold       float64       // Risk score threshold for blocking
	MaxSessionsPerDevice int           // Maximum concurrent sessions per device
}

// SessionInfo tracks active sessions
type SessionInfo struct {
	SessionID   string
	DeviceID    string
	IP          string
	LastActive  time.Time
	UserAgent   string
	RiskScore   float64
	IsBlocked   bool
	BlockReason string
}

// Enhanced DeviceInfo with concurrent access support
type DeviceInfo struct {
	sync.RWMutex
	DeviceID       string
	FirstSeen      time.Time
	LastAccess     time.Time
	VisitCount     int
	IPHistory      []IPRecord
	DeviceProfile  DeviceProfile
	BotStatus      BotDetectionResult
	RiskScore      RiskAssessment
	ActiveSessions map[string]SessionInfo // Track concurrent sessions
	BlockStatus    BlockStatus
}
type IPRecord struct {
	IP        string
	FirstSeen time.Time
	LastSeen  time.Time
	Location  GeoLocation
	Count     int
}

type GeoLocation struct {
	Country     string
	Region      string
	City        string
	ISP         string
	ProxyStatus ProxyAssessment
}

type ProxyAssessment struct {
	IsProxy      bool
	IsVPN        bool
	IsTor        bool
	ThreatLevel  string
	LastDetected time.Time
}

type DeviceProfile struct {
	Hardware     HardwareInfo
	Software     SoftwareInfo
	Capabilities DeviceCapabilities
	LastSeen     time.Time
}

type HardwareInfo struct {
	GPU              string
	CPUCores         int
	Memory           string
	ScreenResolution string
	PixelDensity     float64
	TouchPoints      int
	BatteryStatus    string
	NetworkType      string
	HardwareID       string
}

type SoftwareInfo struct {
	Platform       string
	OS             string
	OSVersion      string
	Browser        string
	BrowserVersion string
	Engine         string
	EngineVersion  string
	Brand          string
	Model          string
	BuildID        string
	Vendor         string
}

type DeviceCapabilities struct {
	WebGL          bool
	Canvas         bool
	WebRTC         bool
	AudioSupport   bool
	VideoSupport   bool
	StorageSupport bool
	SensorsSupport bool
	PWACompatible  bool
}

type BotDetectionResult struct {
	IsBot            bool
	BotType          string
	ConfidenceScore  float64
	DetectionMethod  string
	LastUpdated      time.Time
	BehaviorPatterns []string
}

type RiskAssessment struct {
	Score           float64
	Factors         []RiskFactor
	LastAssessment  time.Time
	RecommendAction string
}

type RiskFactor struct {
	Name        string
	Severity    string
	Description string
	Weight      float64
}

// BlockStatus tracks device blocking information
type BlockStatus struct {
	IsBlocked     bool
	BlockedAt     time.Time
	BlockReason   string
	BlockDuration time.Duration
	BlockHistory  []BlockRecord
}

type BlockRecord struct {
	BlockedAt     time.Time
	UnblockedAt   time.Time
	BlockReason   string
	BlockDuration time.Duration
}

// DeviceStore manages device data with thread-safe operations
type DeviceStore struct {
	sync.RWMutex
	devices   map[string]DeviceInfo
	ipDevices map[string]map[string]bool // IP to DeviceIDs mapping
	limits    DeviceLimits
	store     *session.Store
}
type DeviceAccessResult struct {
	Allowed bool
	Reason  string
}

// Initialize global device store with limits
var deviceStore = NewDeviceStore(DeviceLimits{
	MaxDevicesPerIP:      1,
	MaxIPsPerDevice:      1,
	DeviceTimeout:        24 * time.Hour,
	SuspiciousThreshold:  0.7,
	BlockThreshold:       1.5,
	MaxSessionsPerDevice: 3,
})

func NewDeviceStore(limits DeviceLimits) *DeviceStore {
	return &DeviceStore{
		devices:   make(map[string]DeviceInfo),
		ipDevices: make(map[string]map[string]bool),
		limits:    limits,
		store:     session.New(),
	}
}

// Enhanced fingerprint generation focusing on hardware and system characteristics
func generateEnhancedFingerprint(data map[string]interface{}) string {
	var components []string

	// Hardware-specific components that remain constant across browsers
	if hardware, ok := data["hardware"].(map[string]interface{}); ok {
		components = append(components, fmt.Sprintf("gpu:%v", hardware["gpu"]))
	}
	components = append(components, fmt.Sprintf("cores:%v", data["cores"]))

	if screen, ok := data["screen"].(map[string]interface{}); ok {
		components = append(components, fmt.Sprintf("screen:%vx%v@%v",
			screen["width"], screen["height"], screen["dpi"]))
		if touchPoints, ok := screen["touchPoints"].(float64); ok {
			components = append(components, fmt.Sprintf("touch:%v", touchPoints))
		}
	}

	// System-level identifiers
	if software, ok := data["software"].(map[string]interface{}); ok {
		components = append(components, fmt.Sprintf("platform:%v", software["platform"]))
		components = append(components, fmt.Sprintf("os:%v", parseOS(getString(software, "userAgent"))))
	}

	// Hardware capabilities that are system-dependent

	if sensors, ok := data["sensors"].(map[string]interface{}); ok {
		components = append(components, fmt.Sprintf("accelerometer:%v", getBool(sensors, "accelerometer")))
		components = append(components, fmt.Sprintf("gyroscope:%v", getBool(sensors, "gyroscope")))
	}

	// Memory and performance characteristics
	if memory, ok := data["memory"].(float64); ok {
		components = append(components, fmt.Sprintf("memory:%v", memory))
	}

	// Generate stable hash that should remain consistent across browsers
	fingerprint := strings.Join(components, "|")
	fmt.Printf("Device fingerprint components: %s\n", fingerprint)
	hash := sha256.Sum256([]byte(fingerprint))
	return hex.EncodeToString(hash[:])
}

// Update CheckDeviceAccess to be more strict with device identification
func (ds *DeviceStore) CheckDeviceAccess(deviceID, ip string, data map[string]interface{}) (DeviceAccessResult, error) {
	ds.Lock()
	defer ds.Unlock()

	result := DeviceAccessResult{
		Allowed: true,
		Reason:  "Access granted",
	}

	// Check if this device fingerprint is already blocked
	if device, exists := ds.devices[deviceID]; exists && device.BlockStatus.IsBlocked {
		if time.Since(device.BlockStatus.BlockedAt) < device.BlockStatus.BlockDuration {
			result.Allowed = false
			result.Reason = fmt.Sprintf("Device blocked: %s", device.BlockStatus.BlockReason)
			return result, nil
		}
	}

	// Strict device limit per IP
	if devices, ok := ds.ipDevices[ip]; ok {
		// If this IP has any devices and the current device ID isn't already registered
		if len(devices) > 0 && !devices[deviceID] {
			result.Allowed = false
			result.Reason = fmt.Sprintf("Too many devices from this IP: %d", len(devices))
			return result, nil
		}
	}

	// Get or create device info
	device, exists := ds.devices[deviceID]
	if !exists {
		device = DeviceInfo{
			DeviceID:       deviceID,
			FirstSeen:      time.Now(),
			ActiveSessions: make(map[string]SessionInfo),
		}
	}

	// Update IP tracking
	if !exists {
		if ds.ipDevices[ip] == nil {
			ds.ipDevices[ip] = make(map[string]bool)
		}
		ds.ipDevices[ip][deviceID] = true
	}

	// Enhanced risk assessment
	device.RiskScore = calculateEnhancedRiskScore(device, ip, data)
	if device.RiskScore.Score >= ds.limits.BlockThreshold {
		device.BlockStatus = BlockStatus{
			IsBlocked:     true,
			BlockedAt:     time.Now(),
			BlockReason:   "High risk score detected",
			BlockDuration: 24 * time.Hour,
		}
		result.Allowed = false
		result.Reason = "Access blocked due to suspicious activity"
		ds.devices[deviceID] = device
		return result, nil
	}

	// Update device info
	device.LastAccess = time.Now()
	device.VisitCount++
	ds.devices[deviceID] = device

	return result, nil
}

// Enhanced risk score calculation
func calculateEnhancedRiskScore(device DeviceInfo, currentIP string, data map[string]interface{}) RiskAssessment {
	var factors []RiskFactor
	var totalScore float64

	// Check rapid IP changes
	if len(device.IPHistory) > 0 {
		lastIP := device.IPHistory[len(device.IPHistory)-1]
		timeSinceLastIP := time.Since(lastIP.LastSeen)
		if timeSinceLastIP < 5*time.Minute && lastIP.IP != currentIP {
			factors = append(factors, RiskFactor{
				Name:        "Rapid IP Change",
				Severity:    "High",
				Description: "Multiple IPs in short time period",
				Weight:      0.8,
			})
			totalScore += 0.8
		}
	}

	// Check concurrent sessions
	if len(device.ActiveSessions) > 3 {
		factors = append(factors, RiskFactor{
			Name:        "Multiple Sessions",
			Severity:    "Medium",
			Description: "Excessive concurrent sessions",
			Weight:      0.6,
		})
		totalScore += 0.6
	}

	// Check for bot behavior
	if botScore := detectEnhancedBot(data); botScore > 0 {
		factors = append(factors, RiskFactor{
			Name:        "Bot Behavior",
			Severity:    "High",
			Description: "Automated behavior detected",
			Weight:      botScore,
		})
		totalScore += botScore
	}

	// Check browser fingerprint consistency
	if inconsistencyScore := checkFingerprintConsistency(device, data); inconsistencyScore > 0 {
		factors = append(factors, RiskFactor{
			Name:        "Inconsistent Fingerprint",
			Severity:    "Medium",
			Description: "Device characteristics changed",
			Weight:      inconsistencyScore,
		})
		totalScore += inconsistencyScore
	}

	return RiskAssessment{
		Score:           totalScore,
		Factors:         factors,
		LastAssessment:  time.Now(),
		RecommendAction: getEnhancedRecommendedAction(totalScore),
	}
}

// Enhanced bot detection
func detectEnhancedBot(data map[string]interface{}) float64 {
	var botScore float64

	// Check automation frameworks
	if capabilities, ok := data["capabilities"].(map[string]interface{}); ok {
		if !getBool(capabilities, "webGL") && !getBool(capabilities, "canvas") {
			botScore += 0.4
		}
		if !getBool(capabilities, "audio") || !getBool(capabilities, "video") {
			botScore += 0.3
		}
	}

	// Check user behavior patterns
	if security, ok := data["security"].(map[string]interface{}); ok {
		if !getBool(security, "hasCredentials") {
			botScore += 0.2
		}
	}

	// Check suspicious network patterns
	if network, ok := data["network"].(map[string]interface{}); ok {
		networkType := getString(network, "type")
		if networkType == "" || networkType == "unknown" {
			botScore += 0.3
		}
	}

	return botScore
}

func checkFingerprintConsistency(device DeviceInfo, newData map[string]interface{}) float64 {
	var inconsistencyScore float64

	// Compare with previous profile
	if device.DeviceProfile.LastSeen.IsZero() {
		return 0 // First visit, no comparison possible
	}

	// Check hardware consistency
	if hardware, ok := newData["hardware"].(map[string]interface{}); ok {
		if device.DeviceProfile.Hardware.GPU != getString(hardware, "gpu") {
			inconsistencyScore += 0.4
		}
		if device.DeviceProfile.Hardware.ScreenResolution != getString(hardware, "screenResolution") {
			inconsistencyScore += 0.3
		}
	}

	// Check software consistency
	if software, ok := newData["software"].(map[string]interface{}); ok {
		if device.DeviceProfile.Software.Platform != getString(software, "platform") {
			inconsistencyScore += 0.5
		}
		if device.DeviceProfile.Software.OS != getString(software, "os") {
			inconsistencyScore += 0.5
		}
	}

	return inconsistencyScore
}

func getEnhancedRecommendedAction(score float64) string {
	switch {
	case score > 1.5:
		return "Block"
	case score > 1.0:
		return "Challenge"
	case score > 0.7:
		return "Monitor"
	case score > 0.4:
		return "Verify"
	default:
		return "Allow"
	}
}

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Enhanced Device Detection</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            max-width: 1000px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f0f2f5;
        }
        .container {
            background-color: white;
            padding: 25px;
            border-radius: 12px;
            box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        }
        pre {
            background-color: #f8f9fa;
            padding: 15px;
            border-radius: 8px;
            overflow-x: auto;
            font-family: 'Monaco', 'Consolas', monospace;
        }
        .status {
            margin-top: 20px;
            padding: 15px;
            border-radius: 8px;
            background-color: #e3f2fd;
        }
        .error {
            background-color: #ffebee;
            color: #c62828;
            padding: 10px;
            border-radius: 4px;
            margin: 10px 0;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Enhanced Device Detection</h1>
        <div id="status" class="status">Analyzing device capabilities...</div>
        <div id="error" class="error" style="display: none;"></div>
    </div>

    <script>
        async function getAdvancedFingerprint() {
            const fp = {
                // Basic hardware info
                gpu: await getGPUInfo(),
                cores: navigator.hardwareConcurrency,
                memory: navigator.deviceMemory,
                platform: navigator.platform,
                vendor: navigator.vendor,
                screen: getDetailedScreenInfo(),
                
                // Enhanced device capabilities
                capabilities: await getDeviceCapabilities(),
                
                // Software environment
                software: getDetailedSoftwareInfo(),
                
                // Hardware sensors
                sensors: await getDeviceSensors(),
                
                // Network information
                network: await getNetworkInfo(),
                
                // Additional security metrics
                security: await getSecurityMetrics(),
                
                timestamp: new Date().toISOString()
            };

            return fp;
        }

        async function getGPUInfo() {
            try {
                const canvas = document.createElement('canvas');
                const gl = canvas.getContext('webgl') || canvas.getContext('experimental-webgl');
                if (!gl) return 'WebGL not supported';
                
                const debugInfo = gl.getExtension('WEBGL_debug_renderer_info');
                return debugInfo ? {
                    renderer: gl.getParameter(debugInfo.UNMASKED_RENDERER_WEBGL),
                    vendor: gl.getParameter(debugInfo.UNMASKED_VENDOR_WEBGL),
                    version: gl.getParameter(gl.VERSION)
                } : 'GPU info not available';
            } catch (e) {
                return 'Error getting GPU info';
            }
        }

        function getDetailedScreenInfo() {
            const screen = window.screen;
            return {
                width: screen.width,
                height: screen.height,
                availWidth: screen.availWidth,
                availHeight: screen.availHeight,
                colorDepth: screen.colorDepth,
                pixelDepth: screen.pixelDepth,
                dpi: window.devicePixelRatio,
                orientation: screen.orientation ? screen.orientation.type : 'unknown',
                touchPoints: navigator.maxTouchPoints
            };
        }

        async function getDeviceCapabilities() {
            return {
                webGL: !!document.createElement('canvas').getContext('webgl'),
                webGL2: !!document.createElement('canvas').getContext('webgl2'),
                canvas: !!document.createElement('canvas').getContext('2d'),
                webRTC: !!window.RTCPeerConnection,
                audio: !!document.createElement('audio').canPlayType,
                video: !!document.createElement('video').canPlayType,
                storage: await checkStorageAvailability(),
                pwa: 'serviceWorker' in navigator,
                bluetooth: 'bluetooth' in navigator,
                usb: 'usb' in navigator,
                geolocation: 'geolocation' in navigator
            };
        }

        async function checkStorageAvailability() {
            try {
                return {
                    localStorage: !!window.localStorage,
                    sessionStorage: !!window.sessionStorage,
                    indexedDB: !!window.indexedDB
                };
            } catch {
                return false;
            }
        }

        function getDetailedSoftwareInfo() {
            const ua = navigator.userAgent;
            return {
                platform: navigator.platform,
                userAgent: ua,
                language: navigator.language,
                languages: navigator.languages,
                doNotTrack: navigator.doNotTrack,
                cookieEnabled: navigator.cookieEnabled,
                vendor: navigator.vendor,
                product: navigator.product,
                productSub: navigator.productSub,
                buildID: navigator.buildID || 'unknown'
            };
        }

        async function getDeviceSensors() {
            const sensors = {
                accelerometer: false,
                gyroscope: false,
                magnetometer: false,
                ambient: false
            };

            if ('DeviceMotionEvent' in window) {
                sensors.accelerometer = true;
            }
            if ('DeviceOrientationEvent' in window) {
                sensors.gyroscope = true;
            }
            if ('Magnetometer' in window) {
                sensors.magnetometer = true;
            }
            if ('AmbientLightSensor' in window) {
                sensors.ambient = true;
            }

            return sensors;
        }

        async function getNetworkInfo() {
            const connection = navigator.connection || navigator.mozConnection || navigator.webkitConnection;
            return connection ? {
                type: connection.type,
                effectiveType: connection.effectiveType,
                downlinkMax: connection.downlinkMax,
                rtt: connection.rtt,
                saveData: connection.saveData
            } : 'Network info not available';
        }

        async function getSecurityMetrics() {
            return {
                hasCredentials: !!window.PasswordCredential,
                hasCrypto: !!window.crypto,
                hasSubtleCrypto: !!window.crypto?.subtle,
                securityLevel: await checkSecurityLevel()
            };
        }

        async function checkSecurityLevel() {
            try {
                const hasSecureContext = window.isSecureContext;
                const hasServiceWorker = 'serviceWorker' in navigator;
                const hasHttps = location.protocol === 'https:';
                
                return {
                    secureContext: hasSecureContext,
                    serviceWorker: hasServiceWorker,
                    https: hasHttps
                };
            } catch {
                return 'Security check failed';
            }
        }

        async function sendDeviceInfo() {
            try {
                const deviceData = await getAdvancedFingerprint();
                
                const response = await fetch('/check-device', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(deviceData)
                });

                if (!response.ok) throw new Error('Network response was not ok');
                
                const data = await response.json();
                document.getElementById('status').innerHTML = 
                    '<pre>' + JSON.stringify(data, null, 2) + '</pre>';
                document.getElementById('error').style.display = 'none';
            } catch (error) {
                document.getElementById('error').textContent = 'Error: ' + error.message;
                document.getElementById('error').style.display = 'block';
            }
        }

        sendDeviceInfo();
    </script>
</body>
</html>`

// Main application setup
func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Add middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// Serve HTML template
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString(htmlTemplate) // Using the same template as before
	})

	// Enhanced device check endpoint
	app.Post("/check-device", func(c *fiber.Ctx) error {
		var deviceData map[string]interface{}
		if err := c.BodyParser(&deviceData); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid data format",
			})
		}

		deviceID := generateEnhancedFingerprint(deviceData)
		currentIP := c.IP()

		// Check device access
		accessResult, err := deviceStore.CheckDeviceAccess(deviceID, currentIP, deviceData)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Error checking device access",
			})
		}

		if !accessResult.Allowed {
			return c.Status(403).JSON(fiber.Map{
				"status":  "blocked",
				"reason":  accessResult.Reason,
				"message": "Access denied",
			})
		}

		// Process allowed device
		device, _ := deviceStore.devices[deviceID]
		updateDeviceProfile(&device, deviceData)

		return c.JSON(fiber.Map{
			"status": "success",
			"deviceInfo": fiber.Map{
				"deviceId":   device.DeviceID[:8] + "...",
				"firstSeen":  device.FirstSeen.Format(time.RFC3339),
				"visitCount": device.VisitCount,
				"lastAccess": device.LastAccess.Format(time.RFC3339),
				"profile":    formatDeviceProfile(device.DeviceProfile),
				"riskAssessment": fiber.Map{
					"score":           device.RiskScore.Score,
					"riskLevel":       getRiskLevel(device.RiskScore.Score),
					"factors":         device.RiskScore.Factors,
					"recommendAction": device.RiskScore.RecommendAction,
				},
				"accessStatus": fiber.Map{
					"allowed":        accessResult.Allowed,
					"reason":         accessResult.Reason,
					"activeSessions": len(device.ActiveSessions),
				},
			},
		})
	})

	fmt.Println("Enhanced Device Detection Server starting on http://localhost:3000")
	app.Listen(":3000")
}
func updateDeviceProfile(device *DeviceInfo, data map[string]interface{}) {
	// Update hardware info
	if gpu, ok := data["gpu"].(map[string]interface{}); ok {
		device.DeviceProfile.Hardware.GPU = fmt.Sprintf("%v", gpu["renderer"])
	}
	if cores, ok := data["cores"].(float64); ok {
		device.DeviceProfile.Hardware.CPUCores = int(cores)
	}
	if screen, ok := data["screen"].(map[string]interface{}); ok {
		device.DeviceProfile.Hardware.ScreenResolution = fmt.Sprintf("%vx%v",
			screen["width"], screen["height"])
		if dpi, ok := screen["dpi"].(float64); ok {
			device.DeviceProfile.Hardware.PixelDensity = dpi
		}
		if touchPoints, ok := screen["touchPoints"].(float64); ok {
			device.DeviceProfile.Hardware.TouchPoints = int(touchPoints)
		}
	}

	// Update software info
	if software, ok := data["software"].(map[string]interface{}); ok {
		device.DeviceProfile.Software = SoftwareInfo{
			Platform:       getString(software, "platform"),
			OS:             parseOS(getString(software, "userAgent")),
			OSVersion:      parseOSVersion(getString(software, "userAgent")),
			Browser:        parseBrowser(getString(software, "userAgent")),
			BrowserVersion: parseBrowserVersion(getString(software, "userAgent")),
			Vendor:         getString(software, "vendor"),
			BuildID:        getString(software, "buildID"),
		}
	}

	// Update capabilities
	if capabilities, ok := data["capabilities"].(map[string]interface{}); ok {
		device.DeviceProfile.Capabilities = DeviceCapabilities{
			WebGL:         getBool(capabilities, "webGL"),
			Canvas:        getBool(capabilities, "canvas"),
			WebRTC:        getBool(capabilities, "webRTC"),
			AudioSupport:  getBool(capabilities, "audio"),
			VideoSupport:  getBool(capabilities, "video"),
			PWACompatible: getBool(capabilities, "pwa"),
		}
	}

	device.DeviceProfile.LastSeen = time.Now()
}
func formatDeviceProfile(profile DeviceProfile) fiber.Map {
	return fiber.Map{
		"hardware": fiber.Map{
			"gpu":              profile.Hardware.GPU,
			"cores":            profile.Hardware.CPUCores,
			"screenResolution": profile.Hardware.ScreenResolution,
			"pixelDensity":     profile.Hardware.PixelDensity,
			"touchPoints":      profile.Hardware.TouchPoints,
		},
		"software": fiber.Map{
			"platform": profile.Software.Platform,
			"os":       profile.Software.OS,
			"browser":  profile.Software.Browser,
			"vendor":   profile.Software.Vendor,
		},
		"capabilities": fiber.Map{
			"webGL":  profile.Capabilities.WebGL,
			"canvas": profile.Capabilities.Canvas,
			"webRTC": profile.Capabilities.WebRTC,
			"pwa":    profile.Capabilities.PWACompatible,
		},
	}
}

func getRiskLevel(score float64) string {
	switch {
	case score > 1.5:
		return "High"
	case score > 1.0:
		return "Medium"
	case score > 0.5:
		return "Low"
	default:
		return "Minimal"
	}
}

// Helper functions for type conversion and parsing
func getString(m map[string]interface{}, key string) string {
	if val, ok := m[key].(string); ok {
		return val
	}
	return ""
}
func getBool(m map[string]interface{}, key string) bool {
	if val, ok := m[key].(bool); ok {
		return val
	}
	return false
}

func parseOS(userAgent string) string {
	// Simple OS detection - in production use a proper user agent parser
	userAgent = strings.ToLower(userAgent)
	switch {
	case strings.Contains(userAgent, "windows"):
		return "Windows"
	case strings.Contains(userAgent, "mac"):
		return "macOS"
	case strings.Contains(userAgent, "linux"):
		return "Linux"
	case strings.Contains(userAgent, "android"):
		return "Android"
	case strings.Contains(userAgent, "ios"):
		return "iOS"
	default:
		return "Unknown"
	}
}

func parseOSVersion(userAgent string) string {
	// Implement OS version parsing based on user agent
	return "Version parsing not implemented"
}

func parseBrowser(userAgent string) string {
	// Simple browser detection - in production use a proper user agent parser
	userAgent = strings.ToLower(userAgent)
	switch {
	case strings.Contains(userAgent, "chrome"):
		return "Chrome"
	case strings.Contains(userAgent, "firefox"):
		return "Firefox"
	case strings.Contains(userAgent, "safari"):
		return "Safari"
	case strings.Contains(userAgent, "edge"):
		return "Edge"
	default:
		return "Unknown"
	}
}

func parseBrowserVersion(userAgent string) string {
	// Implement browser version parsing based on user agent
	return "Version parsing not implemented"
}
