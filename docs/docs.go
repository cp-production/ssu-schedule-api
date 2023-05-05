// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/departments": {
            "get": {
                "description": "Retrieves SSU departments' list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "departments"
                ],
                "summary": "get a list of departments",
                "operationId": "get-departments-list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Departments"
                            }
                        }
                    }
                }
            }
        },
        "/{education_form}/{department}/groups": {
            "get": {
                "description": "Retrieves groups' list based on department and education form",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "get a list of groups of a certain department",
                "operationId": "get-groups-list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Education form, e.g. ` + "`" + `do` + "`" + `",
                        "name": "education_form",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Department URL, e.g. ` + "`" + `knt` + "`" + ` for CSIT department",
                        "name": "department",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Groups"
                            }
                        }
                    }
                }
            }
        },
        "/{education_form}/{department}/{group_number}": {
            "get": {
                "description": "Retrieves the schedule based on department, education form and group number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schedule"
                ],
                "summary": "get the schedule of students for a particular group",
                "operationId": "get-students-schedule",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Education form, e.g. ` + "`" + `do` + "`" + `",
                        "name": "education_form",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Department URL, e.g. ` + "`" + `knt` + "`" + ` for CSIT department",
                        "name": "department",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Group number, e.g. ` + "`" + `351` + "`" + `",
                        "name": "group_number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.StudentsLesson"
                            }
                        }
                    }
                }
            }
        },
        "/{education_form}/{department}/{group_number}/subgroups": {
            "get": {
                "description": "Retrieves the subgroups list of a group based on department, education form and group number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "get a list of subgroups of a certain group",
                "operationId": "get-group-subgroups",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Education form, e.g. ` + "`" + `do` + "`" + `",
                        "name": "education_form",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Department URL, e.g. ` + "`" + `knt` + "`" + ` for CSIT department",
                        "name": "department",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Group number, e.g. ` + "`" + `351` + "`" + `",
                        "name": "group_number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Subgroups"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Departments": {
            "type": "object",
            "properties": {
                "fullName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "shortName": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "model.Groups": {
            "type": "object",
            "properties": {
                "depId": {
                    "type": "integer"
                },
                "edForm": {
                    "type": "string"
                },
                "groupNum": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.StudentsLesson": {
            "type": "object",
            "properties": {
                "dayNum": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "lessonName": {
                    "type": "string"
                },
                "lessonPlace": {
                    "type": "string"
                },
                "lessonType": {
                    "type": "string"
                },
                "subgroupName": {
                    "type": "string"
                },
                "teacher": {
                    "type": "string"
                },
                "weekType": {
                    "type": "string"
                }
            }
        },
        "model.Subgroups": {
            "type": "object",
            "properties": {
                "groupId": {
                    "type": "integer"
                },
                "subgroupName": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1.0",
	Schemes:          []string{},
	Title:            "SSU Schedule API",
	Description:      "API Server for SSU Schedule Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
