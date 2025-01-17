package cdktf

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"cdktf.App",
		reflect.TypeOf((*App)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "targetStackId", GoGetter: "TargetStackId"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_App{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.AppOptions",
		reflect.TypeOf((*AppOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.ArtifactoryBackend",
		reflect.TypeOf((*ArtifactoryBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ArtifactoryBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.ArtifactoryBackendProps",
		reflect.TypeOf((*ArtifactoryBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktf.AssetType",
		reflect.TypeOf((*AssetType)(nil)).Elem(),
		map[string]interface{}{
			"FILE": AssetType_FILE,
			"DIRECTORY": AssetType_DIRECTORY,
			"ARCHIVE": AssetType_ARCHIVE,
		},
	)
	_jsii_.RegisterClass(
		"cdktf.AzurermBackend",
		reflect.TypeOf((*AzurermBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_AzurermBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.AzurermBackendProps",
		reflect.TypeOf((*AzurermBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.BooleanMap",
		reflect.TypeOf((*BooleanMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "lookup", GoMethod: "Lookup"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			return &jsiiProxy_BooleanMap{}
		},
	)
	_jsii_.RegisterClass(
		"cdktf.ComplexComputedList",
		reflect.TypeOf((*ComplexComputedList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			return &jsiiProxy_ComplexComputedList{}
		},
	)
	_jsii_.RegisterClass(
		"cdktf.ConsulBackend",
		reflect.TypeOf((*ConsulBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ConsulBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.ConsulBackendProps",
		reflect.TypeOf((*ConsulBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.CosBackend",
		reflect.TypeOf((*CosBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CosBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.CosBackendProps",
		reflect.TypeOf((*CosBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteState",
		reflect.TypeOf((*DataTerraformRemoteState)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteState{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateArtifactory",
		reflect.TypeOf((*DataTerraformRemoteStateArtifactory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateArtifactory{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateArtifactoryConfig",
		reflect.TypeOf((*DataTerraformRemoteStateArtifactoryConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateAzurerm",
		reflect.TypeOf((*DataTerraformRemoteStateAzurerm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateAzurerm{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateAzurermConfig",
		reflect.TypeOf((*DataTerraformRemoteStateAzurermConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateConfig",
		reflect.TypeOf((*DataTerraformRemoteStateConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateConsul",
		reflect.TypeOf((*DataTerraformRemoteStateConsul)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateConsul{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateConsulConfig",
		reflect.TypeOf((*DataTerraformRemoteStateConsulConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateCos",
		reflect.TypeOf((*DataTerraformRemoteStateCos)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateCos{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateCosConfig",
		reflect.TypeOf((*DataTerraformRemoteStateCosConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateEtcd",
		reflect.TypeOf((*DataTerraformRemoteStateEtcd)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateEtcd{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateEtcdConfig",
		reflect.TypeOf((*DataTerraformRemoteStateEtcdConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateEtcdV3",
		reflect.TypeOf((*DataTerraformRemoteStateEtcdV3)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateEtcdV3{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateEtcdV3Config",
		reflect.TypeOf((*DataTerraformRemoteStateEtcdV3Config)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateGcs",
		reflect.TypeOf((*DataTerraformRemoteStateGcs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateGcs{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateGcsConfig",
		reflect.TypeOf((*DataTerraformRemoteStateGcsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateHttp",
		reflect.TypeOf((*DataTerraformRemoteStateHttp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateHttp{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateHttpConfig",
		reflect.TypeOf((*DataTerraformRemoteStateHttpConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateLocal",
		reflect.TypeOf((*DataTerraformRemoteStateLocal)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateLocal{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateLocalConfig",
		reflect.TypeOf((*DataTerraformRemoteStateLocalConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateManta",
		reflect.TypeOf((*DataTerraformRemoteStateManta)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateManta{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateMantaConfig",
		reflect.TypeOf((*DataTerraformRemoteStateMantaConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateOss",
		reflect.TypeOf((*DataTerraformRemoteStateOss)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateOss{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateOssConfig",
		reflect.TypeOf((*DataTerraformRemoteStateOssConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStatePg",
		reflect.TypeOf((*DataTerraformRemoteStatePg)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStatePg{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStatePgConfig",
		reflect.TypeOf((*DataTerraformRemoteStatePgConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateRemoteConfig",
		reflect.TypeOf((*DataTerraformRemoteStateRemoteConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateS3",
		reflect.TypeOf((*DataTerraformRemoteStateS3)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateS3{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateS3Config",
		reflect.TypeOf((*DataTerraformRemoteStateS3Config)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DataTerraformRemoteStateSwift",
		reflect.TypeOf((*DataTerraformRemoteStateSwift)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataTerraformRemoteStateSwift{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformRemoteState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.DataTerraformRemoteStateSwiftConfig",
		reflect.TypeOf((*DataTerraformRemoteStateSwiftConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.DefaultTokenResolver",
		reflect.TypeOf((*DefaultTokenResolver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resolveList", GoMethod: "ResolveList"},
			_jsii_.MemberMethod{JsiiMethod: "resolveString", GoMethod: "ResolveString"},
			_jsii_.MemberMethod{JsiiMethod: "resolveToken", GoMethod: "ResolveToken"},
		},
		func() interface{} {
			j := jsiiProxy_DefaultTokenResolver{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITokenResolver)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.EncodingOptions",
		reflect.TypeOf((*EncodingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.EtcdBackend",
		reflect.TypeOf((*EtcdBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_EtcdBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.EtcdBackendProps",
		reflect.TypeOf((*EtcdBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.EtcdV3Backend",
		reflect.TypeOf((*EtcdV3Backend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_EtcdV3Backend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.EtcdV3BackendProps",
		reflect.TypeOf((*EtcdV3BackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.GcsBackend",
		reflect.TypeOf((*GcsBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_GcsBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.GcsBackendProps",
		reflect.TypeOf((*GcsBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.HttpBackend",
		reflect.TypeOf((*HttpBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_HttpBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.HttpBackendProps",
		reflect.TypeOf((*HttpBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdktf.IAnyProducer",
		reflect.TypeOf((*IAnyProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_IAnyProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.IFragmentConcatenator",
		reflect.TypeOf((*IFragmentConcatenator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "join", GoMethod: "Join"},
		},
		func() interface{} {
			return &jsiiProxy_IFragmentConcatenator{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.IListProducer",
		reflect.TypeOf((*IListProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_IListProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.INumberProducer",
		reflect.TypeOf((*INumberProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_INumberProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.IPostProcessor",
		reflect.TypeOf((*IPostProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "postProcess", GoMethod: "PostProcess"},
		},
		func() interface{} {
			return &jsiiProxy_IPostProcessor{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.IRemoteWorkspace",
		reflect.TypeOf((*IRemoteWorkspace)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IRemoteWorkspace{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.IResolvable",
		reflect.TypeOf((*IResolvable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_IResolvable{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.IResolveContext",
		reflect.TypeOf((*IResolveContext)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "preparing", GoGetter: "Preparing"},
			_jsii_.MemberMethod{JsiiMethod: "registerPostProcessor", GoMethod: "RegisterPostProcessor"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
		},
		func() interface{} {
			return &jsiiProxy_IResolveContext{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.IResource",
		reflect.TypeOf((*IResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IResource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.IStringProducer",
		reflect.TypeOf((*IStringProducer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "produce", GoMethod: "Produce"},
		},
		func() interface{} {
			return &jsiiProxy_IStringProducer{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.ITerraformDependable",
		reflect.TypeOf((*ITerraformDependable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
		},
		func() interface{} {
			return &jsiiProxy_ITerraformDependable{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.ITerraformResource",
		reflect.TypeOf((*ITerraformResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
		},
		func() interface{} {
			return &jsiiProxy_ITerraformResource{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.ITokenMapper",
		reflect.TypeOf((*ITokenMapper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "mapToken", GoMethod: "MapToken"},
		},
		func() interface{} {
			return &jsiiProxy_ITokenMapper{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktf.ITokenResolver",
		reflect.TypeOf((*ITokenResolver)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "resolveList", GoMethod: "ResolveList"},
			_jsii_.MemberMethod{JsiiMethod: "resolveString", GoMethod: "ResolveString"},
			_jsii_.MemberMethod{JsiiMethod: "resolveToken", GoMethod: "ResolveToken"},
		},
		func() interface{} {
			return &jsiiProxy_ITokenResolver{}
		},
	)
	_jsii_.RegisterClass(
		"cdktf.Lazy",
		reflect.TypeOf((*Lazy)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Lazy{}
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.LazyAnyValueOptions",
		reflect.TypeOf((*LazyAnyValueOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktf.LazyListValueOptions",
		reflect.TypeOf((*LazyListValueOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktf.LazyStringValueOptions",
		reflect.TypeOf((*LazyStringValueOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.LocalBackend",
		reflect.TypeOf((*LocalBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_LocalBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.LocalBackendProps",
		reflect.TypeOf((*LocalBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.Manifest",
		reflect.TypeOf((*Manifest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "buildManifest", GoMethod: "BuildManifest"},
			_jsii_.MemberMethod{JsiiMethod: "forStack", GoMethod: "ForStack"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "stacks", GoGetter: "Stacks"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberMethod{JsiiMethod: "writeToFile", GoMethod: "WriteToFile"},
		},
		func() interface{} {
			return &jsiiProxy_Manifest{}
		},
	)
	_jsii_.RegisterClass(
		"cdktf.MantaBackend",
		reflect.TypeOf((*MantaBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_MantaBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.MantaBackendProps",
		reflect.TypeOf((*MantaBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.NamedRemoteWorkspace",
		reflect.TypeOf((*NamedRemoteWorkspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_NamedRemoteWorkspace{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRemoteWorkspace)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdktf.NumberMap",
		reflect.TypeOf((*NumberMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "lookup", GoMethod: "Lookup"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			return &jsiiProxy_NumberMap{}
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.OssAssumeRole",
		reflect.TypeOf((*OssAssumeRole)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.OssBackend",
		reflect.TypeOf((*OssBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_OssBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.OssBackendProps",
		reflect.TypeOf((*OssBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.PgBackend",
		reflect.TypeOf((*PgBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_PgBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.PgBackendProps",
		reflect.TypeOf((*PgBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.PrefixedRemoteWorkspaces",
		reflect.TypeOf((*PrefixedRemoteWorkspaces)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
		},
		func() interface{} {
			j := jsiiProxy_PrefixedRemoteWorkspaces{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRemoteWorkspace)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdktf.RemoteBackend",
		reflect.TypeOf((*RemoteBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_RemoteBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.RemoteBackendProps",
		reflect.TypeOf((*RemoteBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktf.ResolveOptions",
		reflect.TypeOf((*ResolveOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.Resource",
		reflect.TypeOf((*Resource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Resource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdktf.S3Backend",
		reflect.TypeOf((*S3Backend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_S3Backend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.S3BackendProps",
		reflect.TypeOf((*S3BackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktf.StackManifest",
		reflect.TypeOf((*StackManifest)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.StringConcat",
		reflect.TypeOf((*StringConcat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "join", GoMethod: "Join"},
		},
		func() interface{} {
			j := jsiiProxy_StringConcat{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IFragmentConcatenator)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdktf.StringMap",
		reflect.TypeOf((*StringMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "lookup", GoMethod: "Lookup"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			return &jsiiProxy_StringMap{}
		},
	)
	_jsii_.RegisterClass(
		"cdktf.SwiftBackend",
		reflect.TypeOf((*SwiftBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SwiftBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformBackend)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.SwiftBackendProps",
		reflect.TypeOf((*SwiftBackendProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformAsset",
		reflect.TypeOf((*TerraformAsset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assetHash", GoGetter: "AssetHash"},
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformAsset{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformAssetConfig",
		reflect.TypeOf((*TerraformAssetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformBackend",
		reflect.TypeOf((*TerraformBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformBackend{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformElement)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformDataSource",
		reflect.TypeOf((*TerraformDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformElement)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITerraformDependable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITerraformResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformElement",
		reflect.TypeOf((*TerraformElement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformElement{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformElementMetadata",
		reflect.TypeOf((*TerraformElementMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformHclModule",
		reflect.TypeOf((*TerraformHclModule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "set", GoMethod: "Set"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformHclModule{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformModule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformHclModuleOptions",
		reflect.TypeOf((*TerraformHclModuleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformLocal",
		reflect.TypeOf((*TerraformLocal)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "asBoolean", GoGetter: "AsBoolean"},
			_jsii_.MemberProperty{JsiiProperty: "asList", GoGetter: "AsList"},
			_jsii_.MemberProperty{JsiiProperty: "asNumber", GoGetter: "AsNumber"},
			_jsii_.MemberProperty{JsiiProperty: "asString", GoGetter: "AsString"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "expression", GoGetter: "Expression"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformLocal{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformMetaArguments",
		reflect.TypeOf((*TerraformMetaArguments)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformModule",
		reflect.TypeOf((*TerraformModule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformModule{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformElement)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITerraformDependable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformModuleOptions",
		reflect.TypeOf((*TerraformModuleOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformModuleProvider",
		reflect.TypeOf((*TerraformModuleProvider)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformOutput",
		reflect.TypeOf((*TerraformOutput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "sensitive", GoGetter: "Sensitive"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformOutput{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformOutputConfig",
		reflect.TypeOf((*TerraformOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformProvider",
		reflect.TypeOf((*TerraformProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformProvider{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformProviderConfig",
		reflect.TypeOf((*TerraformProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformProviderGeneratorMetadata",
		reflect.TypeOf((*TerraformProviderGeneratorMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformRemoteState",
		reflect.TypeOf((*TerraformRemoteState)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "getBoolean", GoMethod: "GetBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "getList", GoMethod: "GetList"},
			_jsii_.MemberMethod{JsiiMethod: "getNumber", GoMethod: "GetNumber"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformRemoteState{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformElement)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformResource",
		reflect.TypeOf((*TerraformResource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformResource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformElement)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITerraformDependable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformResourceConfig",
		reflect.TypeOf((*TerraformResourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformResourceLifecycle",
		reflect.TypeOf((*TerraformResourceLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformStack",
		reflect.TypeOf((*TerraformStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "allProviders", GoMethod: "AllProviders"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformStack{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformStackMetadata",
		reflect.TypeOf((*TerraformStackMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.TerraformVariable",
		reflect.TypeOf((*TerraformVariable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "booleanValue", GoGetter: "BooleanValue"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNode", GoGetter: "ConstructNode"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "default", GoGetter: "Default"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "listValue", GoGetter: "ListValue"},
			_jsii_.MemberProperty{JsiiProperty: "numberValue", GoGetter: "NumberValue"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "sensitive", GoGetter: "Sensitive"},
			_jsii_.MemberProperty{JsiiProperty: "stringValue", GoGetter: "StringValue"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_TerraformVariable{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_TerraformElement)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktf.TerraformVariableConfig",
		reflect.TypeOf((*TerraformVariableConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktf.Testing",
		reflect.TypeOf((*Testing)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Testing{}
		},
	)
	_jsii_.RegisterClass(
		"cdktf.Token",
		reflect.TypeOf((*Token)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Token{}
		},
	)
	_jsii_.RegisterClass(
		"cdktf.Tokenization",
		reflect.TypeOf((*Tokenization)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Tokenization{}
		},
	)
	_jsii_.RegisterClass(
		"cdktf.TokenizedStringFragments",
		reflect.TypeOf((*TokenizedStringFragments)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIntrinsic", GoMethod: "AddIntrinsic"},
			_jsii_.MemberMethod{JsiiMethod: "addLiteral", GoMethod: "AddLiteral"},
			_jsii_.MemberMethod{JsiiMethod: "addToken", GoMethod: "AddToken"},
			_jsii_.MemberProperty{JsiiProperty: "firstToken", GoGetter: "FirstToken"},
			_jsii_.MemberProperty{JsiiProperty: "firstValue", GoGetter: "FirstValue"},
			_jsii_.MemberMethod{JsiiMethod: "join", GoMethod: "Join"},
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberMethod{JsiiMethod: "mapTokens", GoMethod: "MapTokens"},
			_jsii_.MemberProperty{JsiiProperty: "tokens", GoGetter: "Tokens"},
		},
		func() interface{} {
			return &jsiiProxy_TokenizedStringFragments{}
		},
	)
	_jsii_.RegisterClass(
		"cdktf.VariableType",
		reflect.TypeOf((*VariableType)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_VariableType{}
		},
	)
}
