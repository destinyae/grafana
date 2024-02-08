//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v0alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.Position":            schema_pkg_apis_peakq_v0alpha1_Position(ref),
		"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.QueryTemplate":       schema_pkg_apis_peakq_v0alpha1_QueryTemplate(ref),
		"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.QueryTemplateList":   schema_pkg_apis_peakq_v0alpha1_QueryTemplateList(ref),
		"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.QueryTemplateSpec":   schema_pkg_apis_peakq_v0alpha1_QueryTemplateSpec(ref),
		"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.RenderedQuery":       schema_pkg_apis_peakq_v0alpha1_RenderedQuery(ref),
		"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.Target":              schema_pkg_apis_peakq_v0alpha1_Target(ref),
		"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.TemplateVariable":    schema_pkg_apis_peakq_v0alpha1_TemplateVariable(ref),
		"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.VariableReplacement": schema_pkg_apis_peakq_v0alpha1_VariableReplacement(ref),
	}
}

func schema_pkg_apis_peakq_v0alpha1_Position(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Position is where to do replacement in the targets during render.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"start": {
						SchemaProps: spec.SchemaProps{
							Description: "Start is the byte offset within TargetKey's property of the variable. It is the start location for replacements).",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"end": {
						SchemaProps: spec.SchemaProps{
							Description: "End is the byte offset of the end of the variable.",
							Default:     0,
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
				},
				Required: []string{"start", "end"},
			},
		},
	}
}

func schema_pkg_apis_peakq_v0alpha1_QueryTemplate(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.QueryTemplateSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.QueryTemplateSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_peakq_v0alpha1_QueryTemplateList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.QueryTemplate"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.QueryTemplate", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_peakq_v0alpha1_QueryTemplateSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"title": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"vars": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"key",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "The variables that can be used to render",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.TemplateVariable"),
									},
								},
							},
						},
					},
					"targets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Output variables",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.Target"),
									},
								},
							},
						},
					},
				},
				Required: []string{"title", "targets"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.Target", "github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.TemplateVariable"},
	}
}

func schema_pkg_apis_peakq_v0alpha1_RenderedQuery(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Dummy object that represents a real query object",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"targets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.Target"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.Target"},
	}
}

func schema_pkg_apis_peakq_v0alpha1_Target(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"dataType": {
						SchemaProps: spec.SchemaProps{
							Description: "DataType is the returned Dataplane type from the query.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"variables": {
						SchemaProps: spec.SchemaProps{
							Description: "Variables that will be replaced in the query",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Default: map[string]interface{}{},
													Ref:     ref("github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.VariableReplacement"),
												},
											},
										},
									},
								},
							},
						},
					},
					"properties": {
						SchemaProps: spec.SchemaProps{
							Description: "Query target",
							Ref:         ref("github.com/grafana/grafana/pkg/apis/query/v0alpha1.GenericDataQuery"),
						},
					},
				},
				Required: []string{"variables", "properties"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.VariableReplacement", "github.com/grafana/grafana/pkg/apis/query/v0alpha1.GenericDataQuery"},
	}
}

func schema_pkg_apis_peakq_v0alpha1_TemplateVariable(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TemplateVariable is the definition of a variable that will be interpolated in targets.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"key": {
						SchemaProps: spec.SchemaProps{
							Description: "Key is the name of the variable.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"defaultValue": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "DefaultValue is the value to be used when there is no selected value during render.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"valueListDefinition": {
						SchemaProps: spec.SchemaProps{
							Description: "ValueListDefinition is the object definition used by the FE to get a list of possible values to select for render.",
							Ref:         ref("github.com/grafana/grafana/pkg/apis/common/v0alpha1.Unstructured"),
						},
					},
				},
				Required: []string{"key", "defaultValue", "valueListDefinition"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/common/v0alpha1.Unstructured"},
	}
}

func schema_pkg_apis_peakq_v0alpha1_VariableReplacement(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "QueryVariable is the definition of a variable that will be interpolated in targets.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"path": {
						SchemaProps: spec.SchemaProps{
							Description: "Path is the location of the property within a target. The format for this is not figured out yet (Maybe JSONPath?). Idea: [\"string\", int, \"string\"] where int indicates array offset",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"position": {
						SchemaProps: spec.SchemaProps{
							Description: "Positions is a list of where to perform the interpolation within targets during render. The first string is the Idx of the target as a string, since openAPI does not support ints as map keys",
							Ref:         ref("github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.Position"),
						},
					},
					"format": {
						SchemaProps: spec.SchemaProps{
							Description: "How values should be interpolated\n\nPossible enum values:\n - `\"csv\"` Formats variables with multiple values as a comma-separated string.\n - `\"doublequote\"` Formats single- and multi-valued variables into a comma-separated string\n - `\"json\"` Formats variables with multiple values as a comma-separated string.\n - `\"pipe\"` Formats variables with multiple values into a pipe-separated string.\n - `\"raw\"` Formats variables with multiple values into comma-separated string. This is the default behavior when no format is specified\n - `\"singlequote\"` Formats single- and multi-valued variables into a comma-separated string",
							Type:        []string{"string"},
							Format:      "",
							Enum:        []interface{}{"csv", "doublequote", "json", "pipe", "raw", "singlequote"},
						},
					},
				},
				Required: []string{"path"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/pkg/apis/peakq/v0alpha1.Position"},
	}
}
