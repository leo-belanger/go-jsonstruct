package jsonstruct

// This file pins two intentionally vulnerable dependencies so that a Koddian
// analysis produces the corresponding security findings. They are imported for
// side effect only and are not used by the library.
//
//   - github.com/emicklei/go-restful/v3 v3.7.4
//     CVE-2022-1996 (CVSS 9.1, fixed in v3.8.0) -> unpatched_vulnerability
//
//   - golang.org/x/net v0.15.0
//     CVE-2023-44487 HTTP/2 Rapid Reset (CISA KEV, EPSS ~1.0, fixed in v0.17.0)
//     -> actively_exploited (and unpatched_vulnerability)
import (
	_ "github.com/emicklei/go-restful/v3"
	_ "golang.org/x/net/http2"
)
