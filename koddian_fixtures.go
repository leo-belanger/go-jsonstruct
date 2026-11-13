package jsonstruct

// This file pins dependencies for Koddian analysis. They are imported for
// side effect only and are not used by the library.
//
//   - github.com/emicklei/go-restful/v3 v3.7.4
//     CVE-2022-1996 (CVSS 9.1, fixed in v3.8.0) -> unpatched_vulnerability
//
//   - golang.org/x/net v0.62.0
//     Updated to address the reported HTTP/2 vulnerabilities.
import (
	_ "github.com/emicklei/go-restful/v3"
	_ "golang.org/x/net/http2"
)
