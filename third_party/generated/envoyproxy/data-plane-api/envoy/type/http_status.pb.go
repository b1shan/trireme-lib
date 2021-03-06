// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/http_status.proto

package envoy_type

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// HTTP response codes supported in Envoy.
// For more details: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type StatusCode int32

const (
	// Empty - This code not part of the HTTP status code specification, but it is needed for proto
	// `enum` type.
	StatusCode_Empty                         StatusCode = 0
	StatusCode_Continue                      StatusCode = 100
	StatusCode_OK                            StatusCode = 200
	StatusCode_Created                       StatusCode = 201
	StatusCode_Accepted                      StatusCode = 202
	StatusCode_NonAuthoritativeInformation   StatusCode = 203
	StatusCode_NoContent                     StatusCode = 204
	StatusCode_ResetContent                  StatusCode = 205
	StatusCode_PartialContent                StatusCode = 206
	StatusCode_MultiStatus                   StatusCode = 207
	StatusCode_AlreadyReported               StatusCode = 208
	StatusCode_IMUsed                        StatusCode = 226
	StatusCode_MultipleChoices               StatusCode = 300
	StatusCode_MovedPermanently              StatusCode = 301
	StatusCode_Found                         StatusCode = 302
	StatusCode_SeeOther                      StatusCode = 303
	StatusCode_NotModified                   StatusCode = 304
	StatusCode_UseProxy                      StatusCode = 305
	StatusCode_TemporaryRedirect             StatusCode = 307
	StatusCode_PermanentRedirect             StatusCode = 308
	StatusCode_BadRequest                    StatusCode = 400
	StatusCode_Unauthorized                  StatusCode = 401
	StatusCode_PaymentRequired               StatusCode = 402
	StatusCode_Forbidden                     StatusCode = 403
	StatusCode_NotFound                      StatusCode = 404
	StatusCode_MethodNotAllowed              StatusCode = 405
	StatusCode_NotAcceptable                 StatusCode = 406
	StatusCode_ProxyAuthenticationRequired   StatusCode = 407
	StatusCode_RequestTimeout                StatusCode = 408
	StatusCode_Conflict                      StatusCode = 409
	StatusCode_Gone                          StatusCode = 410
	StatusCode_LengthRequired                StatusCode = 411
	StatusCode_PreconditionFailed            StatusCode = 412
	StatusCode_PayloadTooLarge               StatusCode = 413
	StatusCode_URITooLong                    StatusCode = 414
	StatusCode_UnsupportedMediaType          StatusCode = 415
	StatusCode_RangeNotSatisfiable           StatusCode = 416
	StatusCode_ExpectationFailed             StatusCode = 417
	StatusCode_MisdirectedRequest            StatusCode = 421
	StatusCode_UnprocessableEntity           StatusCode = 422
	StatusCode_Locked                        StatusCode = 423
	StatusCode_FailedDependency              StatusCode = 424
	StatusCode_UpgradeRequired               StatusCode = 426
	StatusCode_PreconditionRequired          StatusCode = 428
	StatusCode_TooManyRequests               StatusCode = 429
	StatusCode_RequestHeaderFieldsTooLarge   StatusCode = 431
	StatusCode_InternalServerError           StatusCode = 500
	StatusCode_NotImplemented                StatusCode = 501
	StatusCode_BadGateway                    StatusCode = 502
	StatusCode_ServiceUnavailable            StatusCode = 503
	StatusCode_GatewayTimeout                StatusCode = 504
	StatusCode_HTTPVersionNotSupported       StatusCode = 505
	StatusCode_VariantAlsoNegotiates         StatusCode = 506
	StatusCode_InsufficientStorage           StatusCode = 507
	StatusCode_LoopDetected                  StatusCode = 508
	StatusCode_NotExtended                   StatusCode = 510
	StatusCode_NetworkAuthenticationRequired StatusCode = 511
)

var StatusCode_name = map[int32]string{
	0:   "Empty",
	100: "Continue",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "NonAuthoritativeInformation",
	204: "NoContent",
	205: "ResetContent",
	206: "PartialContent",
	207: "MultiStatus",
	208: "AlreadyReported",
	226: "IMUsed",
	300: "MultipleChoices",
	301: "MovedPermanently",
	302: "Found",
	303: "SeeOther",
	304: "NotModified",
	305: "UseProxy",
	307: "TemporaryRedirect",
	308: "PermanentRedirect",
	400: "BadRequest",
	401: "Unauthorized",
	402: "PaymentRequired",
	403: "Forbidden",
	404: "NotFound",
	405: "MethodNotAllowed",
	406: "NotAcceptable",
	407: "ProxyAuthenticationRequired",
	408: "RequestTimeout",
	409: "Conflict",
	410: "Gone",
	411: "LengthRequired",
	412: "PreconditionFailed",
	413: "PayloadTooLarge",
	414: "URITooLong",
	415: "UnsupportedMediaType",
	416: "RangeNotSatisfiable",
	417: "ExpectationFailed",
	421: "MisdirectedRequest",
	422: "UnprocessableEntity",
	423: "Locked",
	424: "FailedDependency",
	426: "UpgradeRequired",
	428: "PreconditionRequired",
	429: "TooManyRequests",
	431: "RequestHeaderFieldsTooLarge",
	500: "InternalServerError",
	501: "NotImplemented",
	502: "BadGateway",
	503: "ServiceUnavailable",
	504: "GatewayTimeout",
	505: "HTTPVersionNotSupported",
	506: "VariantAlsoNegotiates",
	507: "InsufficientStorage",
	508: "LoopDetected",
	510: "NotExtended",
	511: "NetworkAuthenticationRequired",
}

var StatusCode_value = map[string]int32{
	"Empty":                         0,
	"Continue":                      100,
	"OK":                            200,
	"Created":                       201,
	"Accepted":                      202,
	"NonAuthoritativeInformation":   203,
	"NoContent":                     204,
	"ResetContent":                  205,
	"PartialContent":                206,
	"MultiStatus":                   207,
	"AlreadyReported":               208,
	"IMUsed":                        226,
	"MultipleChoices":               300,
	"MovedPermanently":              301,
	"Found":                         302,
	"SeeOther":                      303,
	"NotModified":                   304,
	"UseProxy":                      305,
	"TemporaryRedirect":             307,
	"PermanentRedirect":             308,
	"BadRequest":                    400,
	"Unauthorized":                  401,
	"PaymentRequired":               402,
	"Forbidden":                     403,
	"NotFound":                      404,
	"MethodNotAllowed":              405,
	"NotAcceptable":                 406,
	"ProxyAuthenticationRequired":   407,
	"RequestTimeout":                408,
	"Conflict":                      409,
	"Gone":                          410,
	"LengthRequired":                411,
	"PreconditionFailed":            412,
	"PayloadTooLarge":               413,
	"URITooLong":                    414,
	"UnsupportedMediaType":          415,
	"RangeNotSatisfiable":           416,
	"ExpectationFailed":             417,
	"MisdirectedRequest":            421,
	"UnprocessableEntity":           422,
	"Locked":                        423,
	"FailedDependency":              424,
	"UpgradeRequired":               426,
	"PreconditionRequired":          428,
	"TooManyRequests":               429,
	"RequestHeaderFieldsTooLarge":   431,
	"InternalServerError":           500,
	"NotImplemented":                501,
	"BadGateway":                    502,
	"ServiceUnavailable":            503,
	"GatewayTimeout":                504,
	"HTTPVersionNotSupported":       505,
	"VariantAlsoNegotiates":         506,
	"InsufficientStorage":           507,
	"LoopDetected":                  508,
	"NotExtended":                   510,
	"NetworkAuthenticationRequired": 511,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}

func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7544d7adacd3389b, []int{0}
}

// HTTP status.
type HttpStatus struct {
	// Supplies HTTP response code.
	Code                 StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.type.StatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HttpStatus) Reset()         { *m = HttpStatus{} }
func (m *HttpStatus) String() string { return proto.CompactTextString(m) }
func (*HttpStatus) ProtoMessage()    {}
func (*HttpStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_7544d7adacd3389b, []int{0}
}
func (m *HttpStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpStatus.Merge(m, src)
}
func (m *HttpStatus) XXX_Size() int {
	return m.Size()
}
func (m *HttpStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HttpStatus proto.InternalMessageInfo

func (m *HttpStatus) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_Empty
}

func init() {
	proto.RegisterEnum("envoy.type.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterType((*HttpStatus)(nil), "envoy.type.HttpStatus")
}

func init() { proto.RegisterFile("envoy/type/http_status.proto", fileDescriptor_7544d7adacd3389b) }

var fileDescriptor_7544d7adacd3389b = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x49, 0x6f, 0x5c, 0xc5,
	0x13, 0xcf, 0x9b, 0xce, 0xe6, 0x4e, 0x62, 0x57, 0x3a, 0x8b, 0xfd, 0xcf, 0x3f, 0x58, 0x56, 0x4e,
	0x88, 0x83, 0x2d, 0xc1, 0x09, 0x89, 0x8b, 0xed, 0xd8, 0xb1, 0xc1, 0x33, 0x19, 0x8d, 0x67, 0x72,
	0x45, 0xed, 0xd7, 0x35, 0x33, 0xad, 0xbc, 0xe9, 0x7a, 0xe9, 0x57, 0x33, 0xf6, 0xe3, 0xc8, 0x27,
	0x60, 0xdf, 0xd7, 0x03, 0x8b, 0x50, 0x42, 0x40, 0xc0, 0x77, 0x08, 0x7b, 0x3e, 0x02, 0xf2, 0x67,
	0x60, 0x0d, 0x08, 0x50, 0xf7, 0x2c, 0xf6, 0x85, 0x93, 0xfd, 0xaa, 0x6b, 0xf9, 0x2d, 0x35, 0x25,
	0x2f, 0xa3, 0x1b, 0x50, 0xb9, 0xc4, 0x65, 0x8e, 0x4b, 0x5d, 0xe6, 0xfc, 0xe9, 0x82, 0x35, 0xf7,
	0x8b, 0xc5, 0xdc, 0x13, 0x93, 0x92, 0xf1, 0x75, 0x31, 0xbc, 0x5e, 0x9a, 0x1d, 0xe8, 0xcc, 0x1a,
	0xcd, 0xb8, 0x34, 0xfe, 0x67, 0x98, 0x74, 0xe5, 0x49, 0x29, 0x37, 0x98, 0xf3, 0xed, 0x58, 0xa8,
	0x9e, 0x90, 0x47, 0x53, 0x32, 0x38, 0x97, 0x2c, 0x24, 0x0f, 0x4f, 0x3f, 0x7a, 0x71, 0xf1, 0xa0,
	0xc3, 0xe2, 0x30, 0x63, 0x95, 0x0c, 0xae, 0xc0, 0x83, 0x95, 0x63, 0xcf, 0x26, 0x95, 0x85, 0x23,
	0xc3, 0xbf, 0x90, 0x34, 0x62, 0xd5, 0x23, 0x5f, 0x4d, 0x49, 0x79, 0x90, 0xa6, 0xa6, 0xe4, 0xb1,
	0xb5, 0x5e, 0xce, 0x25, 0x1c, 0x51, 0xa7, 0xe5, 0xc9, 0x55, 0x72, 0x6c, 0x5d, 0x1f, 0xc1, 0xa8,
	0x13, 0xb2, 0x72, 0xfd, 0x29, 0xb8, 0x97, 0xa8, 0xd3, 0xf2, 0xc4, 0xaa, 0x47, 0xcd, 0x68, 0xe0,
	0xeb, 0x44, 0x9d, 0x91, 0x27, 0x97, 0xd3, 0x14, 0xf3, 0xf0, 0xf9, 0x4d, 0xa2, 0x16, 0xe4, 0xff,
	0x6b, 0xe4, 0x96, 0xfb, 0xdc, 0x25, 0x6f, 0x59, 0xb3, 0x1d, 0xe0, 0xa6, 0x6b, 0x93, 0xef, 0x69,
	0xb6, 0xe4, 0xe0, 0xdb, 0x44, 0x4d, 0xcb, 0xa9, 0x1a, 0x85, 0xbe, 0xe8, 0x18, 0xbe, 0x4b, 0xd4,
	0x59, 0x79, 0xba, 0x81, 0x05, 0xf2, 0x38, 0xf4, 0x7d, 0xa2, 0xce, 0xc9, 0xe9, 0xba, 0xf6, 0x6c,
	0x75, 0x36, 0x0e, 0xfe, 0x90, 0x28, 0x90, 0xa7, 0xaa, 0xfd, 0x8c, 0xed, 0x10, 0x2b, 0xfc, 0x98,
	0xa8, 0xf3, 0x72, 0x66, 0x39, 0xf3, 0xa8, 0x4d, 0xd9, 0xc0, 0x9c, 0x7c, 0x40, 0x70, 0x3f, 0x51,
	0xa7, 0xe4, 0xf1, 0xcd, 0x6a, 0xab, 0x40, 0x03, 0xfb, 0x31, 0x25, 0x16, 0xe5, 0x19, 0xae, 0x76,
	0xc9, 0xa6, 0x58, 0xc0, 0xed, 0x8a, 0xba, 0x20, 0xa1, 0x4a, 0x03, 0x34, 0x75, 0xf4, 0x3d, 0xed,
	0xd0, 0x71, 0x56, 0xc2, 0x9d, 0x8a, 0x92, 0xf2, 0xd8, 0x3a, 0xf5, 0x9d, 0x81, 0x4f, 0x2b, 0x81,
	0xd6, 0x36, 0xe2, 0x75, 0xee, 0xa2, 0x87, 0xbb, 0x95, 0x30, 0xbc, 0x46, 0x5c, 0x25, 0x63, 0xdb,
	0x16, 0x0d, 0x7c, 0x16, 0x13, 0x5a, 0x05, 0xd6, 0x3d, 0xed, 0x95, 0xf0, 0x79, 0x45, 0x5d, 0x94,
	0x67, 0x9b, 0xd8, 0xcb, 0xc9, 0x6b, 0x5f, 0x36, 0xd0, 0x58, 0x8f, 0x29, 0xc3, 0x17, 0x31, 0x3e,
	0x99, 0x32, 0x89, 0x7f, 0x59, 0x51, 0x33, 0x52, 0xae, 0x68, 0xd3, 0xc0, 0x5b, 0x7d, 0x2c, 0x18,
	0x9e, 0x13, 0x41, 0x86, 0x96, 0xd3, 0x43, 0xdd, 0x9e, 0x41, 0x03, 0xcf, 0x8b, 0x00, 0xbe, 0xae,
	0xcb, 0x5e, 0xac, 0xbc, 0xd5, 0xb7, 0x1e, 0x0d, 0xbc, 0x20, 0x82, 0x7e, 0xeb, 0xe4, 0x77, 0xac,
	0x31, 0xe8, 0xe0, 0x45, 0x11, 0x80, 0xd4, 0x88, 0x87, 0xc0, 0x5f, 0x12, 0x91, 0x1b, 0x72, 0x97,
	0x4c, 0x8d, 0x78, 0x39, 0xcb, 0x68, 0x17, 0x0d, 0xbc, 0x2c, 0x94, 0x92, 0x67, 0x42, 0x20, 0x3a,
	0xa5, 0x77, 0x32, 0x84, 0x57, 0x44, 0xf0, 0x2a, 0xe2, 0x0f, 0x6e, 0xa1, 0x63, 0x9b, 0x46, 0x8f,
	0x26, 0xb3, 0x5e, 0x15, 0xc1, 0x88, 0x11, 0xc4, 0xa6, 0xed, 0x21, 0xf5, 0x19, 0x5e, 0x8b, 0x03,
	0x57, 0xc9, 0xb5, 0x33, 0x9b, 0x32, 0xbc, 0x2e, 0xd4, 0x94, 0x3c, 0x7a, 0x8d, 0x1c, 0xc2, 0x1b,
	0x31, 0x7d, 0x0b, 0x5d, 0x87, 0xbb, 0x93, 0x1e, 0x6f, 0x0a, 0x35, 0x2b, 0x55, 0xdd, 0x63, 0x4a,
	0xce, 0xd8, 0xd0, 0x7e, 0x5d, 0xdb, 0x0c, 0x0d, 0xbc, 0x35, 0xa6, 0x97, 0x91, 0x36, 0x4d, 0xa2,
	0x2d, 0xed, 0x3b, 0x08, 0x6f, 0x8b, 0x20, 0x4c, 0xab, 0xb1, 0x19, 0x22, 0xe4, 0x3a, 0xf0, 0x8e,
	0x50, 0xff, 0x93, 0xe7, 0x5b, 0xae, 0xe8, 0xe7, 0x43, 0x87, 0xab, 0x68, 0xac, 0x6e, 0x96, 0x39,
	0xc2, 0xbb, 0x42, 0xcd, 0xc9, 0x73, 0x0d, 0xed, 0x3a, 0x58, 0x23, 0xde, 0xd6, 0x6c, 0x8b, 0xb6,
	0x8d, 0xd4, 0xde, 0x13, 0x41, 0xf6, 0xb5, 0xbd, 0x1c, 0x53, 0xd6, 0x87, 0x66, 0xbe, 0x1f, 0xc1,
	0x54, 0x6d, 0x31, 0xb4, 0x01, 0x27, 0xf2, 0x7f, 0x10, 0x5b, 0xb5, 0x5c, 0xee, 0x29, 0xc5, 0xa2,
	0x08, 0x4d, 0xd6, 0x1c, 0x5b, 0x2e, 0xe1, 0x43, 0x11, 0xf6, 0x69, 0x8b, 0xd2, 0x9b, 0x68, 0xe0,
	0xa3, 0xa8, 0xee, 0xb0, 0xd9, 0x55, 0xcc, 0xd1, 0x19, 0x74, 0x69, 0x09, 0x1f, 0x47, 0x2a, 0xad,
	0xbc, 0xe3, 0xb5, 0xc1, 0x09, 0xf3, 0x4f, 0x22, 0xf2, 0xc3, 0xcc, 0x27, 0x4f, 0xb7, 0x63, 0x41,
	0x93, 0xa8, 0xaa, 0x5d, 0x39, 0xc2, 0x50, 0xc0, 0x9d, 0x68, 0xc8, 0xe8, 0x73, 0x03, 0xb5, 0x41,
	0xbf, 0x6e, 0x31, 0x33, 0xc5, 0x44, 0x9d, 0xbb, 0x11, 0xe6, 0xa6, 0x63, 0xf4, 0x4e, 0x67, 0xdb,
	0xe8, 0x07, 0xe8, 0xd7, 0xbc, 0x27, 0x0f, 0x3f, 0x47, 0xed, 0x6b, 0xc4, 0x9b, 0xbd, 0x3c, 0xc3,
	0xb0, 0x31, 0x68, 0xe0, 0x17, 0x31, 0xda, 0xb2, 0x6b, 0x9a, 0x71, 0x57, 0x97, 0xf0, 0x6b, 0xe4,
	0x1f, 0xea, 0x6c, 0x8a, 0x2d, 0xa7, 0x07, 0xda, 0x66, 0x51, 0xb0, 0xdf, 0x62, 0xf9, 0x28, 0x6d,
	0xec, 0xf4, 0xef, 0x42, 0x5d, 0x96, 0xb3, 0x1b, 0xcd, 0x66, 0xfd, 0x06, 0xfa, 0xc2, 0x92, 0x0b,
	0x2a, 0x8f, 0x6d, 0x80, 0x3f, 0x84, 0xba, 0x24, 0x2f, 0xdc, 0xd0, 0xde, 0x6a, 0xc7, 0xcb, 0x59,
	0x41, 0x35, 0xec, 0x10, 0x5b, 0xcd, 0x58, 0xc0, 0x83, 0x11, 0xce, 0xa2, 0xdf, 0x6e, 0xdb, 0xd4,
	0xa2, 0xe3, 0x6d, 0x26, 0xaf, 0x3b, 0x08, 0x7f, 0xc6, 0x3d, 0xdf, 0x22, 0xca, 0xaf, 0x22, 0x47,
	0x0b, 0xe0, 0x2f, 0x31, 0xfa, 0x71, 0xad, 0xed, 0x71, 0x50, 0xd4, 0xc0, 0xdf, 0x42, 0x5d, 0x91,
	0x0f, 0xd5, 0x90, 0x77, 0xc9, 0xdf, 0xfc, 0x8f, 0xdd, 0xfc, 0x47, 0xac, 0x3c, 0x7e, 0x6f, 0x7f,
	0x3e, 0xb9, 0xbf, 0x3f, 0x9f, 0xfc, 0xb4, 0x3f, 0x9f, 0xc8, 0x39, 0x4b, 0xc3, 0xbb, 0x97, 0x87,
	0x8d, 0x3e, 0x74, 0x02, 0x57, 0x66, 0x0e, 0x2e, 0x65, 0x3d, 0x1c, 0xcf, 0x7a, 0xb2, 0x73, 0x3c,
	0x5e, 0xd1, 0xc7, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x49, 0xbb, 0xde, 0x8a, 0x05, 0x00,
	0x00,
}

func (m *HttpStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHttpStatus(dAtA, i, uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintHttpStatus(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HttpStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovHttpStatus(uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHttpStatus(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHttpStatus(x uint64) (n int) {
	return sovHttpStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HttpStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= StatusCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHttpStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttpStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHttpStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttpStatus
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHttpStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHttpStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthHttpStatus
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHttpStatus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttpStatus
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHttpStatus(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHttpStatus
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHttpStatus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttpStatus   = fmt.Errorf("proto: integer overflow")
)
