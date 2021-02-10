package godaddy

type RecordType int

const (
    RecordTypeA          = 1  // a host address
    RecordTypeNS         = 2  // an authoritative name server
    RecordTypeMD         = 3  // a mail destination (OBSOLETE - use MX)
    RecordTypeMF         = 4  // a mail forwarder (OBSOLETE - use MX)
    RecordTypeCNAME      = 5  // the canonical name for an alias
    RecordTypeSOA        = 6  // marks the start of a zone of authority
    RecordTypeMB         = 7  // a mailbox domain name (EXPERIMENTAL)
    RecordTypeMG         = 8  // a mail group member (EXPERIMENTAL)
    RecordTypeMR         = 9  // a mail rename domain name (EXPERIMENTAL)
    RecordTypeNULL       = 10 // a null RR (EXPERIMENTAL)
    RecordTypeWKS        = 11 // a well known service description
    RecordTypePTR        = 12 // a domain name pointer
    RecordTypeHINFO      = 13 // host information
    RecordTypeMINFO      = 14 // mailbox or mail list information
    RecordTypeMX         = 15 // mail exchange
    RecordTypeTXT        = 16 // text strings
    RecordTypeRP         = 17 // for Responsible Person
    RecordTypeAFSDB      = 18 // for AFS Data Base location
    RecordTypeX25        = 19 // for X.25 PSDN address
    RecordTypeISDN       = 20 // for ISDN address
    RecordTypeRT         = 21 // for Route Through
    RecordTypeNSAP       = 22 // for NSAP address, NSAP style A record
    RecordTypeNSAP_PTR   = 23 // spelled "NSAP-PTR", for domain name pointer, NSAP style
    RecordTypeSIG        = 24 // for security signature
    RecordTypeKEY        = 25 // for security key
    RecordTypePX         = 26 // X.400 mail mapping information
    RecordTypeGPOS       = 27 // Geographical Position
    RecordTypeAAAA       = 28 // IP6 Address
    RecordTypeLOC        = 29 // Location Information
    RecordTypeNXT        = 30 // Next Domain (OBSOLETE)
    RecordTypeEID        = 31 // Endpoint Identifier
    RecordTypeNIMLOC     = 32 // Nimrod Locator
    RecordTypeSRV        = 33 // Server Selection
    RecordTypeATMA       = 34 // ATM Address
    RecordTypeNAPTR      = 35 // Naming Authority Pointer
    RecordTypeKX         = 36 // Key Exchanger
    RecordTypeCERT       = 37 // CERT
    RecordTypeA6         = 38 // A6 (OBSOLETE - use AAAA)
    RecordTypeDNAME      = 39 // DNAME
    RecordTypeSINK       = 40 // SINK
    RecordTypeOPT        = 41 // OPT
    RecordTypeAPL        = 42 // APL
    RecordTypeDS         = 43 // Delegation Signer
    RecordTypeSSHFP      = 44 // SSH Key Fingerprint
    RecordTypeIPSECKEY   = 45 // IPSECKEY
    RecordTypeRRSIG      = 46 // RRSIG
    RecordTypeNSEC       = 47 // NSEC
    RecordTypeDNSKEY     = 48 // DNSKEY
    RecordTypeDHCID      = 49 // DHCID
    RecordTypeNSEC3      = 50 // NSEC3
    RecordTypeNSEC3PARAM = 51 // NSEC3PARAM
    RecordTypeTLSA       = 52 // TLSA
    RecordTypeSMIMEA     = 53 // S/MIME cert association

    RecordTypeHIP        = 55 // Host Identity Protocol
    RecordTypeNINFO      = 56 // NINFO
    RecordTypeRKEY       = 57 // RKEY
    RecordTypeTALINK     = 58 // Trust Anchor LINK
    RecordTypeCDS        = 59 // Child DS
    RecordTypeCDNSKEY    = 60 // DNSKEY(s) the Child wants reflected in DS
    RecordTypeOPENPGPKEY = 61 // OpenPGP Key
    RecordTypeCSYNC      = 62 // Child-To-Parent Synchronization
    RecordTypeZONEMD     = 63 // message digest for DNS zone

    RecordTypeSPF    = 99  // declares which hosts are, and are not, authorized to use a domain name for the "HELO" and "MAIL FROM" identities (OBSOLETE - use TXT)
    RecordTypeUINFO  = 100 // [IANA-Reserved]
    RecordTypeUID    = 101 // [IANA-Reserved]
    RecordTypeGID    = 102 // [IANA-Reserved]
    RecordTypeUNSPEC = 103 // [IANA-Reserved]
    RecordTypeNID    = 104 // values for Node Identifiers that will be used for ILNP-capable nodes
    RecordTypeL32    = 105 // 32-bit Locator values for ILNPv4-capable nodes
    RecordTypeL64    = 106 // unsigned 64-bit Locator values for ILNPv6-capable nodes
    RecordTypeLP     = 107 // the name of a subnetwork for ILNP
    RecordTypeEUI48  = 108 // an EUI-48 address
    RecordTypeEUI64  = 109 // an EUI-64 address

    RecordTypeTKEY     = 249 // Transaction Key
    RecordTypeTSIG     = 250 // Transaction Signature
    RecordTypeIXFR     = 251 // incremental transfer
    RecordTypeAXFR     = 252 // transfer of an entire zone
    RecordTypeMAILB    = 253 // mailbox-related RRs (MB, MG or MR)
    RecordTypeMAILA    = 254 // mail agent RRs (OBSOLETE - see MX)
    RecordTypeAll      = 255 // Spelled "*", A request for some or all records the server has available
    RecordTypeURI      = 256 // URI
    RecordTypeCAA      = 257 // Certification Authority Restriction
    RecordTypeAVC      = 258 // Application Visibility and Control
    RecordTypeDOA      = 259 // Digital Object Architecture
    RecordTypeAMTRELAY = 260 // Automatic Multicast Tunneling Relay

    RecordTypeTA  = 32768 // DNSSEC Trust Authorities
    RecordTypeDLV = 32769 // DNSSEC Lookaside Validation

    // Unassigned	32770-65279
    // Private use	65280-65534
    // Reserved	65535
)

var RecordTypeMap = map[string]int{
    "A":          RecordTypeA,
    "NS":         RecordTypeNS,
    "MD":         RecordTypeMD,
    "MF":         RecordTypeMF,
    "CNAME":      RecordTypeCNAME,
    "SOA":        RecordTypeSOA,
    "MB":         RecordTypeMB,
    "MG":         RecordTypeMG,
    "MR":         RecordTypeMR,
    "NULL":       RecordTypeNULL,
    "WKS":        RecordTypeWKS,
    "PTR":        RecordTypePTR,
    "HINFO":      RecordTypeHINFO,
    "MINFO":      RecordTypeMINFO,
    "MX":         RecordTypeMX,
    "TXT":        RecordTypeTXT,
    "RP":         RecordTypeRP,
    "AFSDB":      RecordTypeAFSDB,
    "X25":        RecordTypeX25,
    "ISDN":       RecordTypeISDN,
    "RT":         RecordTypeRT,
    "NSAP":       RecordTypeNSAP,
    "NSAP-PTR":   RecordTypeNSAP_PTR,
    "SIG":        RecordTypeSIG,
    "KEY":        RecordTypeKEY,
    "PX":         RecordTypePX,
    "GPOS":       RecordTypeGPOS,
    "AAAA":       RecordTypeAAAA,
    "LOC":        RecordTypeLOC,
    "NXT":        RecordTypeNXT,
    "EID":        RecordTypeEID,
    "NIMLOC":     RecordTypeNIMLOC,
    "SRV":        RecordTypeSRV,
    "ATMA":       RecordTypeATMA,
    "NAPTR":      RecordTypeNAPTR,
    "KX":         RecordTypeKX,
    "CERT":       RecordTypeCERT,
    "A6":         RecordTypeA6,
    "DNAME":      RecordTypeDNAME,
    "SINK":       RecordTypeSINK,
    "OPT":        RecordTypeOPT,
    "APL":        RecordTypeAPL,
    "DS":         RecordTypeDS,
    "SSHFP":      RecordTypeSSHFP,
    "IPSECKEY":   RecordTypeIPSECKEY,
    "RRSIG":      RecordTypeRRSIG,
    "NSEC":       RecordTypeNSEC,
    "DNSKEY":     RecordTypeDNSKEY,
    "DHCID":      RecordTypeDHCID,
    "NSEC3":      RecordTypeNSEC3,
    "NSEC3PARAM": RecordTypeNSEC3PARAM,
    "TLSA":       RecordTypeTLSA,
    "SMIMEA":     RecordTypeSMIMEA,
    "HIP":        RecordTypeHIP,
    "NINFO":      RecordTypeNINFO,
    "RKEY":       RecordTypeRKEY,
    "TALINK":     RecordTypeTALINK,
    "CDS":        RecordTypeCDS,
    "CDNSKEY":    RecordTypeCDNSKEY,
    "OPENPGPKEY": RecordTypeOPENPGPKEY,
    "CSYNC":      RecordTypeCSYNC,
    "ZONEMD":     RecordTypeZONEMD,
    "SPF":        RecordTypeSPF,
    "UINFO":      RecordTypeUINFO,
    "UID":        RecordTypeUID,
    "GID":        RecordTypeGID,
    "UNSPEC":     RecordTypeUNSPEC,
    "NID":        RecordTypeNID,
    "L32":        RecordTypeL32,
    "L64":        RecordTypeL64,
    "LP":         RecordTypeLP,
    "EUI48":      RecordTypeEUI48,
    "EUI64":      RecordTypeEUI64,
    "TKEY":       RecordTypeTKEY,
    "TSIG":       RecordTypeTSIG,
    "IXFR":       RecordTypeIXFR,
    "AXFR":       RecordTypeAXFR,
    "MAILB":      RecordTypeMAILB,
    "MAILA":      RecordTypeMAILA,
    "*":          RecordTypeAll,
    "URI":        RecordTypeURI,
    "CAA":        RecordTypeCAA,
    "AVC":        RecordTypeAVC,
    "DOA":        RecordTypeDOA,
    "AMTRELAY":   RecordTypeAMTRELAY,
    "TA":         RecordTypeTA,
    "DLV":        RecordTypeDLV,
}
