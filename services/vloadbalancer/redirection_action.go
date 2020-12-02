/*
 * vloadbalancer
 *
 * VPC Load Balancer 관련 API<br/>https://ncloud.apigw.ntruss.com/vloadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancer

type RedirectionAction struct {

	// 프로토콜
Protocol *string `json:"protocol,omitempty"`

	// 포트
Port *string `json:"port,omitempty"`

	// 호스트
Host *string `json:"host,omitempty"`

	// 경로
Path *string `json:"path,omitempty"`

	// 쿼리
Query *string `json:"query,omitempty"`

	// 상태코드
StatusCode *string `json:"statusCode,omitempty"`
}
