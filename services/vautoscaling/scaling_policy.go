/*
 * vautoscaling
 *
 * VPC Auto Scaling 관련 API<br/>https://ncloud.apigw.ntruss.com/vautoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vautoscaling

type ScalingPolicy struct {

	// 오토스케일링그룹번호
AutoScalingGroupNo *string `json:"autoScalingGroupNo,omitempty"`

	// 정책번호
PolicyNo *string `json:"policyNo,omitempty"`

	// 정책이름
PolicyName *string `json:"policyName,omitempty"`

	// 조정유형
AdjustmentType *CommonCode `json:"adjustmentType,omitempty"`

	// 조정값
ScalingAdjustment *int32 `json:"scalingAdjustment,omitempty"`

	// 최소조정폭
MinAdjustmentStep *int32 `json:"minAdjustmentStep,omitempty"`

	// 쿨다운
CoolDown *int32 `json:"coolDown,omitempty"`
}
