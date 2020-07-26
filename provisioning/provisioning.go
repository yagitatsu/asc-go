package provisioning

import "github.com/aaronsky/asc-go/internal/services"

// Service handles communication with provisioning-related methods of the App Store Connect API
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundle_ids
// https://developer.apple.com/documentation/appstoreconnectapi/bundle_id_capabilities
// https://developer.apple.com/documentation/appstoreconnectapi/certificates
// https://developer.apple.com/documentation/appstoreconnectapi/devices
// https://developer.apple.com/documentation/appstoreconnectapi/profiles
type Service services.Service