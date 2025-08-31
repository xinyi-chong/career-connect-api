package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// Certificate
type GetCertificatesReq struct {
	g.Meta `path:"/certificates" tags:"User" method:"get" summary:"Get Certificates"`
}

type GetCertificatesRes struct {
	Certificates []*entity.Certificate `json:"certificates"`
}

type GetCertificatesByUserIDReq struct {
	g.Meta `path:"/:user_id/certificates" tags:"User" method:"get" summary:"Get Certificates By User ID"`
	UserID string `json:"user_id"     v:"required"`
}

type GetCertificatesByUserIDRes struct {
	Certificates []*entity.Certificate `json:"certificates"`
}

type GetCertificateByIDReq struct {
	g.Meta `path:"/certificate/:certificate_id" tags:"User" method:"get" summary:"Get Certificate By ID"`
}

type GetCertificateByIDRes struct {
	Certificate *entity.Certificate `json:"certificate"`
}

type PostCreateCertificateReq struct {
	g.Meta `path:"/certificate" tags:"User" method:"post" summary:"Create Certificate"`
	Name   string `json:"name"     v:"required"`
}

type PostCreateCertificateRes struct {
	Id string `json:"id"`
}

type PatchUpdateCertificateByIDReq struct {
	g.Meta `path:"/certificate/:certificate_id" tags:"User" method:"patch" summary:"Update Certificate By ID"`
	Name   string `json:"name"     v:"required"`
}

type PatchUpdateCertificateByIDRes struct {
}

type DeleteCertificateByIDReq struct {
	g.Meta `path:"/certificate/:certificate_id" tags:"User" method:"delete" summary:"Delete Certificate By ID"`
}

type DeleteCertificateByIDRes struct {
}

// CertService
type GetCertificateCertServiceReq struct {
	g.Meta `path:"/certificate/certservice" tags:"User" method:"get" summary:"Get First Cert"`
}

type GetCertificateCertServiceRes struct {
	CertServiceCertificates *map[string]interface{}
}

// type PostCreateCertificateCertServiceReq struct {
// 	g.Meta `path:"/certificate/certservice" tags:"User" method:"post" summary:"Create First Cert"`

// }

// type PostCreateCertificateCertServiceRes struct {
// 	Id     string      `json:"id"`
// }
