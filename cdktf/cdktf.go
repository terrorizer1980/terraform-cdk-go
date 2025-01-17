// Cloud Development Kit for Terraform
package cdktf

import (
	_jsii_ "github.com/aws/jsii-runtime-go"
	_init_ "github.com/hashicorp/terraform-cdk-go/cdktf/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/hashicorp/terraform-cdk-go/cdktf/internal"
)

// Represents a cdktf application.
// Experimental.
type App interface {
	constructs.Construct
	Outdir() *string
	TargetStackId() *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Synth()
	ToString() *string
}

// The jsii proxy struct for App
type jsiiProxy_App struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_App) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) TargetStackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStackId",
		&returns,
	)
	return returns
}


// Defines an app.
// Experimental.
func NewApp(options *AppOptions) App {
	_init_.Initialize()

	j := jsiiProxy_App{}

	_jsii_.Create(
		"cdktf.App",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Defines an app.
// Experimental.
func NewApp_Override(a App, options *AppOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.App",
		[]interface{}{options},
		a,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_App) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_App) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (a *jsiiProxy_App) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Synthesizes all resources to the output directory.
// Experimental.
func (a *jsiiProxy_App) Synth() {
	_jsii_.InvokeVoid(
		a,
		"synth",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_App) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type AppOptions struct {
	// Additional context values for the application.
	//
	// Context set by the CLI or the `context` key in `cdktf.json` has precedence.
	//
	// Context can be read from any construct using `node.getContext(key)`.
	// Experimental.
	Context *map[string]interface{} `json:"context"`
	// The directory to output Terraform resources.
	// Experimental.
	Outdir *string `json:"outdir"`
	// Experimental.
	StackTraces *bool `json:"stackTraces"`
}

// Experimental.
type ArtifactoryBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ArtifactoryBackend
type jsiiProxy_ArtifactoryBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_ArtifactoryBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactoryBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewArtifactoryBackend(scope constructs.Construct, props *ArtifactoryBackendProps) ArtifactoryBackend {
	_init_.Initialize()

	j := jsiiProxy_ArtifactoryBackend{}

	_jsii_.Create(
		"cdktf.ArtifactoryBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewArtifactoryBackend_Override(a ArtifactoryBackend, scope constructs.Construct, props *ArtifactoryBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ArtifactoryBackend",
		[]interface{}{scope, props},
		a,
	)
}

// Experimental.
func (a *jsiiProxy_ArtifactoryBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_ArtifactoryBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_ArtifactoryBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (a *jsiiProxy_ArtifactoryBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_ArtifactoryBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (a *jsiiProxy_ArtifactoryBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ArtifactoryBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_ArtifactoryBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ArtifactoryBackendProps struct {
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Repo *string `json:"repo"`
	// Experimental.
	Subpath *string `json:"subpath"`
	// Experimental.
	Url *string `json:"url"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type AssetType string

const (
	AssetType_FILE AssetType = "FILE"
	AssetType_DIRECTORY AssetType = "DIRECTORY"
	AssetType_ARCHIVE AssetType = "ARCHIVE"
)

// Experimental.
type AzurermBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AzurermBackend
type jsiiProxy_AzurermBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_AzurermBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewAzurermBackend(scope constructs.Construct, props *AzurermBackendProps) AzurermBackend {
	_init_.Initialize()

	j := jsiiProxy_AzurermBackend{}

	_jsii_.Create(
		"cdktf.AzurermBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAzurermBackend_Override(a AzurermBackend, scope constructs.Construct, props *AzurermBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.AzurermBackend",
		[]interface{}{scope, props},
		a,
	)
}

// Experimental.
func (a *jsiiProxy_AzurermBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AzurermBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AzurermBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (a *jsiiProxy_AzurermBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AzurermBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (a *jsiiProxy_AzurermBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AzurermBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AzurermBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type AzurermBackendProps struct {
	// Experimental.
	ContainerName *string `json:"containerName"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	StorageAccountName *string `json:"storageAccountName"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	ClientId *string `json:"clientId"`
	// Experimental.
	ClientSecret *string `json:"clientSecret"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	Environment *string `json:"environment"`
	// Experimental.
	MsiEndpoint *string `json:"msiEndpoint"`
	// Experimental.
	ResourceGroupName *string `json:"resourceGroupName"`
	// Experimental.
	SasToken *string `json:"sasToken"`
	// Experimental.
	SubscriptionId *string `json:"subscriptionId"`
	// Experimental.
	TenantId *string `json:"tenantId"`
	// Experimental.
	UseMsi *bool `json:"useMsi"`
}

// Experimental.
type BooleanMap interface {
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() ITerraformResource
	SetTerraformResource(val ITerraformResource)
	Lookup(key *string) *bool
}

// The jsii proxy struct for BooleanMap
type jsiiProxy_BooleanMap struct {
	_ byte // padding
}

func (j *jsiiProxy_BooleanMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BooleanMap) TerraformResource() ITerraformResource {
	var returns ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewBooleanMap(terraformResource ITerraformResource, terraformAttribute *string) BooleanMap {
	_init_.Initialize()

	j := jsiiProxy_BooleanMap{}

	_jsii_.Create(
		"cdktf.BooleanMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewBooleanMap_Override(b BooleanMap, terraformResource ITerraformResource, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.BooleanMap",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BooleanMap) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BooleanMap) SetTerraformResource(val ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BooleanMap) Lookup(key *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		b,
		"lookup",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Experimental.
type ComplexComputedList interface {
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() ITerraformResource
	SetTerraformResource(val ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) *bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) *string
}

// The jsii proxy struct for ComplexComputedList
type jsiiProxy_ComplexComputedList struct {
	_ byte // padding
}

func (j *jsiiProxy_ComplexComputedList) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexComputedList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComplexComputedList) TerraformResource() ITerraformResource {
	var returns ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewComplexComputedList(terraformResource ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) ComplexComputedList {
	_init_.Initialize()

	j := jsiiProxy_ComplexComputedList{}

	_jsii_.Create(
		"cdktf.ComplexComputedList",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewComplexComputedList_Override(c ComplexComputedList, terraformResource ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ComplexComputedList",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		c,
	)
}

func (j *jsiiProxy_ComplexComputedList) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_ComplexComputedList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComplexComputedList) SetTerraformResource(val ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_ComplexComputedList) GetBooleanAttribute(terraformAttribute *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ComplexComputedList) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ComplexComputedList) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ComplexComputedList) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_ComplexComputedList) InterpolationForAttribute(property *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Experimental.
type ConsulBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ConsulBackend
type jsiiProxy_ConsulBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_ConsulBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConsulBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewConsulBackend(scope constructs.Construct, props *ConsulBackendProps) ConsulBackend {
	_init_.Initialize()

	j := jsiiProxy_ConsulBackend{}

	_jsii_.Create(
		"cdktf.ConsulBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewConsulBackend_Override(c ConsulBackend, scope constructs.Construct, props *ConsulBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.ConsulBackend",
		[]interface{}{scope, props},
		c,
	)
}

// Experimental.
func (c *jsiiProxy_ConsulBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_ConsulBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_ConsulBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (c *jsiiProxy_ConsulBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_ConsulBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (c *jsiiProxy_ConsulBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_ConsulBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_ConsulBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ConsulBackendProps struct {
	// Experimental.
	AccessToken *string `json:"accessToken"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	Address *string `json:"address"`
	// Experimental.
	CaFile *string `json:"caFile"`
	// Experimental.
	CertFile *string `json:"certFile"`
	// Experimental.
	Datacenter *string `json:"datacenter"`
	// Experimental.
	Gzip *bool `json:"gzip"`
	// Experimental.
	HttpAuth *string `json:"httpAuth"`
	// Experimental.
	KeyFile *string `json:"keyFile"`
	// Experimental.
	Lock *bool `json:"lock"`
	// Experimental.
	Scheme *string `json:"scheme"`
}

// Experimental.
type CosBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CosBackend
type jsiiProxy_CosBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_CosBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewCosBackend(scope constructs.Construct, props *CosBackendProps) CosBackend {
	_init_.Initialize()

	j := jsiiProxy_CosBackend{}

	_jsii_.Create(
		"cdktf.CosBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCosBackend_Override(c CosBackend, scope constructs.Construct, props *CosBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.CosBackend",
		[]interface{}{scope, props},
		c,
	)
}

// Experimental.
func (c *jsiiProxy_CosBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CosBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CosBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (c *jsiiProxy_CosBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CosBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (c *jsiiProxy_CosBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CosBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CosBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type CosBackendProps struct {
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	SecretId *string `json:"secretId"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
}

// Experimental.
type DataTerraformRemoteState interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteState
type jsiiProxy_DataTerraformRemoteState struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteState) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteState) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteState) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteState) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteState) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteState(scope constructs.Construct, id *string, config *DataTerraformRemoteStateRemoteConfig) DataTerraformRemoteState {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteState{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteState",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteState_Override(d DataTerraformRemoteState, scope constructs.Construct, id *string, config *DataTerraformRemoteStateRemoteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteState",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteState) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteState) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateArtifactory interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateArtifactory
type jsiiProxy_DataTerraformRemoteStateArtifactory struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateArtifactory) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateArtifactory(scope constructs.Construct, id *string, config *DataTerraformRemoteStateArtifactoryConfig) DataTerraformRemoteStateArtifactory {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateArtifactory{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateArtifactory",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateArtifactory_Override(d DataTerraformRemoteStateArtifactory, scope constructs.Construct, id *string, config *DataTerraformRemoteStateArtifactoryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateArtifactory",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateArtifactory) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateArtifactoryConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Repo *string `json:"repo"`
	// Experimental.
	Subpath *string `json:"subpath"`
	// Experimental.
	Url *string `json:"url"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type DataTerraformRemoteStateAzurerm interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateAzurerm
type jsiiProxy_DataTerraformRemoteStateAzurerm struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateAzurerm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateAzurerm(scope constructs.Construct, id *string, config *DataTerraformRemoteStateAzurermConfig) DataTerraformRemoteStateAzurerm {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateAzurerm{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateAzurerm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateAzurerm_Override(d DataTerraformRemoteStateAzurerm, scope constructs.Construct, id *string, config *DataTerraformRemoteStateAzurermConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateAzurerm",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateAzurerm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateAzurermConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	ContainerName *string `json:"containerName"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	StorageAccountName *string `json:"storageAccountName"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	ClientId *string `json:"clientId"`
	// Experimental.
	ClientSecret *string `json:"clientSecret"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	Environment *string `json:"environment"`
	// Experimental.
	MsiEndpoint *string `json:"msiEndpoint"`
	// Experimental.
	ResourceGroupName *string `json:"resourceGroupName"`
	// Experimental.
	SasToken *string `json:"sasToken"`
	// Experimental.
	SubscriptionId *string `json:"subscriptionId"`
	// Experimental.
	TenantId *string `json:"tenantId"`
	// Experimental.
	UseMsi *bool `json:"useMsi"`
}

// Experimental.
type DataTerraformRemoteStateConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
}

// Experimental.
type DataTerraformRemoteStateConsul interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateConsul
type jsiiProxy_DataTerraformRemoteStateConsul struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateConsul) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateConsul(scope constructs.Construct, id *string, config *DataTerraformRemoteStateConsulConfig) DataTerraformRemoteStateConsul {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateConsul{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateConsul",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateConsul_Override(d DataTerraformRemoteStateConsul, scope constructs.Construct, id *string, config *DataTerraformRemoteStateConsulConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateConsul",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateConsul) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateConsulConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	AccessToken *string `json:"accessToken"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	Address *string `json:"address"`
	// Experimental.
	CaFile *string `json:"caFile"`
	// Experimental.
	CertFile *string `json:"certFile"`
	// Experimental.
	Datacenter *string `json:"datacenter"`
	// Experimental.
	Gzip *bool `json:"gzip"`
	// Experimental.
	HttpAuth *string `json:"httpAuth"`
	// Experimental.
	KeyFile *string `json:"keyFile"`
	// Experimental.
	Lock *bool `json:"lock"`
	// Experimental.
	Scheme *string `json:"scheme"`
}

// Experimental.
type DataTerraformRemoteStateCos interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateCos
type jsiiProxy_DataTerraformRemoteStateCos struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateCos) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateCos(scope constructs.Construct, id *string, config *DataTerraformRemoteStateCosConfig) DataTerraformRemoteStateCos {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateCos{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateCos",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateCos_Override(d DataTerraformRemoteStateCos, scope constructs.Construct, id *string, config *DataTerraformRemoteStateCosConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateCos",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateCos) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateCos) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateCosConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	SecretId *string `json:"secretId"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
}

// Experimental.
type DataTerraformRemoteStateEtcd interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateEtcd
type jsiiProxy_DataTerraformRemoteStateEtcd struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcd) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateEtcd(scope constructs.Construct, id *string, config *DataTerraformRemoteStateEtcdConfig) DataTerraformRemoteStateEtcd {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateEtcd{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateEtcd",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateEtcd_Override(d DataTerraformRemoteStateEtcd, scope constructs.Construct, id *string, config *DataTerraformRemoteStateEtcdConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateEtcd",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcd) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateEtcdConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Endpoints *string `json:"endpoints"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type DataTerraformRemoteStateEtcdV3 interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateEtcdV3
type jsiiProxy_DataTerraformRemoteStateEtcdV3 struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateEtcdV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateEtcdV3(scope constructs.Construct, id *string, config *DataTerraformRemoteStateEtcdV3Config) DataTerraformRemoteStateEtcdV3 {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateEtcdV3{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateEtcdV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateEtcdV3_Override(d DataTerraformRemoteStateEtcdV3, scope constructs.Construct, id *string, config *DataTerraformRemoteStateEtcdV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateEtcdV3",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateEtcdV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateEtcdV3Config struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Endpoints *[]*string `json:"endpoints"`
	// Experimental.
	CacertPath *string `json:"cacertPath"`
	// Experimental.
	CertPath *string `json:"certPath"`
	// Experimental.
	KeyPath *string `json:"keyPath"`
	// Experimental.
	Lock *bool `json:"lock"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type DataTerraformRemoteStateGcs interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateGcs
type jsiiProxy_DataTerraformRemoteStateGcs struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateGcs) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateGcs(scope constructs.Construct, id *string, config *DataTerraformRemoteStateGcsConfig) DataTerraformRemoteStateGcs {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateGcs{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateGcs",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateGcs_Override(d DataTerraformRemoteStateGcs, scope constructs.Construct, id *string, config *DataTerraformRemoteStateGcsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateGcs",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateGcs) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateGcsConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	AccessToken *string `json:"accessToken"`
	// Experimental.
	Credentials *string `json:"credentials"`
	// Experimental.
	EncryptionKey *string `json:"encryptionKey"`
	// Experimental.
	Prefix *string `json:"prefix"`
}

// Experimental.
type DataTerraformRemoteStateHttp interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateHttp
type jsiiProxy_DataTerraformRemoteStateHttp struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateHttp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateHttp(scope constructs.Construct, id *string, config *DataTerraformRemoteStateHttpConfig) DataTerraformRemoteStateHttp {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateHttp{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateHttp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateHttp_Override(d DataTerraformRemoteStateHttp, scope constructs.Construct, id *string, config *DataTerraformRemoteStateHttpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateHttp",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateHttp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateHttpConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Address *string `json:"address"`
	// Experimental.
	LockAddress *string `json:"lockAddress"`
	// Experimental.
	LockMethod *string `json:"lockMethod"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	RetryMax *float64 `json:"retryMax"`
	// Experimental.
	RetryWaitMax *float64 `json:"retryWaitMax"`
	// Experimental.
	RetryWaitMin *float64 `json:"retryWaitMin"`
	// Experimental.
	SkipCertVerification *bool `json:"skipCertVerification"`
	// Experimental.
	UnlockAddress *string `json:"unlockAddress"`
	// Experimental.
	UnlockMethod *string `json:"unlockMethod"`
	// Experimental.
	UpdateMethod *string `json:"updateMethod"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type DataTerraformRemoteStateLocal interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateLocal
type jsiiProxy_DataTerraformRemoteStateLocal struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateLocal) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateLocal(scope constructs.Construct, id *string, config *DataTerraformRemoteStateLocalConfig) DataTerraformRemoteStateLocal {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateLocal{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateLocal",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateLocal_Override(d DataTerraformRemoteStateLocal, scope constructs.Construct, id *string, config *DataTerraformRemoteStateLocalConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateLocal",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateLocal) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateLocalConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	WorkspaceDir *string `json:"workspaceDir"`
}

// Experimental.
type DataTerraformRemoteStateManta interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateManta
type jsiiProxy_DataTerraformRemoteStateManta struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateManta) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateManta(scope constructs.Construct, id *string, config *DataTerraformRemoteStateMantaConfig) DataTerraformRemoteStateManta {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateManta{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateManta",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateManta_Override(d DataTerraformRemoteStateManta, scope constructs.Construct, id *string, config *DataTerraformRemoteStateMantaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateManta",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateManta) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateManta) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateMantaConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Account *string `json:"account"`
	// Experimental.
	KeyId *string `json:"keyId"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	InsecureSkipTlsVerify *bool `json:"insecureSkipTlsVerify"`
	// Experimental.
	KeyMaterial *string `json:"keyMaterial"`
	// Experimental.
	ObjectName *string `json:"objectName"`
	// Experimental.
	Url *string `json:"url"`
	// Experimental.
	User *string `json:"user"`
}

// Experimental.
type DataTerraformRemoteStateOss interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateOss
type jsiiProxy_DataTerraformRemoteStateOss struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateOss) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateOss(scope constructs.Construct, id *string, config *DataTerraformRemoteStateOssConfig) DataTerraformRemoteStateOss {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateOss{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateOss",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateOss_Override(d DataTerraformRemoteStateOss, scope constructs.Construct, id *string, config *DataTerraformRemoteStateOssConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateOss",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateOss) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateOss) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateOssConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	AssumeRole *OssAssumeRole `json:"assumeRole"`
	// Experimental.
	EcsRoleName *string `json:"ecsRoleName"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Profile *string `json:"profile"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
	// Experimental.
	SecurityToken *string `json:"securityToken"`
	// Experimental.
	SharedCredentialsFile *string `json:"sharedCredentialsFile"`
	// Experimental.
	TablestoreEndpoint *string `json:"tablestoreEndpoint"`
	// Experimental.
	TablestoreTable *string `json:"tablestoreTable"`
}

// Experimental.
type DataTerraformRemoteStatePg interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStatePg
type jsiiProxy_DataTerraformRemoteStatePg struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStatePg) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStatePg(scope constructs.Construct, id *string, config *DataTerraformRemoteStatePgConfig) DataTerraformRemoteStatePg {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStatePg{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStatePg",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStatePg_Override(d DataTerraformRemoteStatePg, scope constructs.Construct, id *string, config *DataTerraformRemoteStatePgConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStatePg",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStatePg) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStatePg) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStatePgConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	ConnStr *string `json:"connStr"`
	// Experimental.
	SchemaName *string `json:"schemaName"`
	// Experimental.
	SkipSchemaCreation *bool `json:"skipSchemaCreation"`
}

// Experimental.
type DataTerraformRemoteStateRemoteConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Organization *string `json:"organization"`
	// Experimental.
	Workspaces IRemoteWorkspace `json:"workspaces"`
	// Experimental.
	Hostname *string `json:"hostname"`
	// Experimental.
	Token *string `json:"token"`
}

// Experimental.
type DataTerraformRemoteStateS3 interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateS3
type jsiiProxy_DataTerraformRemoteStateS3 struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateS3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateS3(scope constructs.Construct, id *string, config *DataTerraformRemoteStateS3Config) DataTerraformRemoteStateS3 {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateS3{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateS3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateS3_Override(d DataTerraformRemoteStateS3, scope constructs.Construct, id *string, config *DataTerraformRemoteStateS3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateS3",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateS3) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateS3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateS3Config struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	AssumeRolePolicy *string `json:"assumeRolePolicy"`
	// Experimental.
	DynamodbEndpoint *string `json:"dynamodbEndpoint"`
	// Experimental.
	DynamodbTable *string `json:"dynamodbTable"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	ExternalId *string `json:"externalId"`
	// Experimental.
	ForcePathStyle *bool `json:"forcePathStyle"`
	// Experimental.
	IamEndpoint *string `json:"iamEndpoint"`
	// Experimental.
	KmsKeyId *string `json:"kmsKeyId"`
	// Experimental.
	MaxRetries *float64 `json:"maxRetries"`
	// Experimental.
	Profile *string `json:"profile"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	RoleArn *string `json:"roleArn"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
	// Experimental.
	SessionName *string `json:"sessionName"`
	// Experimental.
	SharedCredentialsFile *string `json:"sharedCredentialsFile"`
	// Experimental.
	SkipCredentialsValidation *bool `json:"skipCredentialsValidation"`
	// Experimental.
	SkipMetadataApiCheck *bool `json:"skipMetadataApiCheck"`
	// Experimental.
	SseCustomerKey *string `json:"sseCustomerKey"`
	// Experimental.
	StsEndpoint *string `json:"stsEndpoint"`
	// Experimental.
	Token *string `json:"token"`
	// Experimental.
	WorkspaceKeyPrefix *string `json:"workspaceKeyPrefix"`
}

// Experimental.
type DataTerraformRemoteStateSwift interface {
	TerraformRemoteState
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataTerraformRemoteStateSwift
type jsiiProxy_DataTerraformRemoteStateSwift struct {
	jsiiProxy_TerraformRemoteState
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTerraformRemoteStateSwift) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTerraformRemoteStateSwift(scope constructs.Construct, id *string, config *DataTerraformRemoteStateSwiftConfig) DataTerraformRemoteStateSwift {
	_init_.Initialize()

	j := jsiiProxy_DataTerraformRemoteStateSwift{}

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateSwift",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTerraformRemoteStateSwift_Override(d DataTerraformRemoteStateSwift, scope constructs.Construct, id *string, config *DataTerraformRemoteStateSwiftConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DataTerraformRemoteStateSwift",
		[]interface{}{scope, id, config},
		d,
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		d,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataTerraformRemoteStateSwift) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type DataTerraformRemoteStateSwiftConfig struct {
	// Experimental.
	Defaults *map[string]interface{} `json:"defaults"`
	// Experimental.
	Workspace *string `json:"workspace"`
	// Experimental.
	Container *string `json:"container"`
	// Experimental.
	ApplicationCredentialId *string `json:"applicationCredentialId"`
	// Experimental.
	ApplicationCredentialName *string `json:"applicationCredentialName"`
	// Experimental.
	ApplicationCredentialSecret *string `json:"applicationCredentialSecret"`
	// Experimental.
	ArchiveContainer *string `json:"archiveContainer"`
	// Experimental.
	AuthUrl *string `json:"authUrl"`
	// Experimental.
	CacertFile *string `json:"cacertFile"`
	// Experimental.
	Cert *string `json:"cert"`
	// Experimental.
	Cloud *string `json:"cloud"`
	// Experimental.
	DefaultDomain *string `json:"defaultDomain"`
	// Experimental.
	DomainId *string `json:"domainId"`
	// Experimental.
	DomainName *string `json:"domainName"`
	// Experimental.
	ExpireAfter *string `json:"expireAfter"`
	// Experimental.
	Insecure *bool `json:"insecure"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	ProjectDomainId *string `json:"projectDomainId"`
	// Experimental.
	ProjectDomainName *string `json:"projectDomainName"`
	// Experimental.
	RegionName *string `json:"regionName"`
	// Experimental.
	StateName *string `json:"stateName"`
	// Experimental.
	TenantId *string `json:"tenantId"`
	// Experimental.
	TenantName *string `json:"tenantName"`
	// Experimental.
	Token *string `json:"token"`
	// Experimental.
	UserDomainId *string `json:"userDomainId"`
	// Experimental.
	UserDomainName *string `json:"userDomainName"`
	// Experimental.
	UserId *string `json:"userId"`
	// Experimental.
	UserName *string `json:"userName"`
}

// Default resolver implementation.
// Experimental.
type DefaultTokenResolver interface {
	ITokenResolver
	ResolveList(xs *[]*string, context IResolveContext) interface{}
	ResolveString(fragments TokenizedStringFragments, context IResolveContext) interface{}
	ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{}
}

// The jsii proxy struct for DefaultTokenResolver
type jsiiProxy_DefaultTokenResolver struct {
	jsiiProxy_ITokenResolver
}

// Experimental.
func NewDefaultTokenResolver(concat IFragmentConcatenator) DefaultTokenResolver {
	_init_.Initialize()

	j := jsiiProxy_DefaultTokenResolver{}

	_jsii_.Create(
		"cdktf.DefaultTokenResolver",
		[]interface{}{concat},
		&j,
	)

	return &j
}

// Experimental.
func NewDefaultTokenResolver_Override(d DefaultTokenResolver, concat IFragmentConcatenator) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.DefaultTokenResolver",
		[]interface{}{concat},
		d,
	)
}

// Resolve a tokenized list.
// Experimental.
func (d *jsiiProxy_DefaultTokenResolver) ResolveList(xs *[]*string, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveList",
		[]interface{}{xs, context},
		&returns,
	)

	return returns
}

// Resolve string fragments to Tokens.
// Experimental.
func (d *jsiiProxy_DefaultTokenResolver) ResolveString(fragments TokenizedStringFragments, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveString",
		[]interface{}{fragments, context},
		&returns,
	)

	return returns
}

// Default Token resolution.
//
// Resolve the Token, recurse into whatever it returns,
// then finally post-process it.
// Experimental.
func (d *jsiiProxy_DefaultTokenResolver) ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolveToken",
		[]interface{}{t, context, postProcessor},
		&returns,
	)

	return returns
}

// Properties to string encodings.
// Experimental.
type EncodingOptions struct {
	// A hint for the Token's purpose when stringifying it.
	// Experimental.
	DisplayHint *string `json:"displayHint"`
}

// Experimental.
type EtcdBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for EtcdBackend
type jsiiProxy_EtcdBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_EtcdBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewEtcdBackend(scope constructs.Construct, props *EtcdBackendProps) EtcdBackend {
	_init_.Initialize()

	j := jsiiProxy_EtcdBackend{}

	_jsii_.Create(
		"cdktf.EtcdBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEtcdBackend_Override(e EtcdBackend, scope constructs.Construct, props *EtcdBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.EtcdBackend",
		[]interface{}{scope, props},
		e,
	)
}

// Experimental.
func (e *jsiiProxy_EtcdBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EtcdBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EtcdBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (e *jsiiProxy_EtcdBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EtcdBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (e *jsiiProxy_EtcdBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EtcdBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_EtcdBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type EtcdBackendProps struct {
	// Experimental.
	Endpoints *string `json:"endpoints"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type EtcdV3Backend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for EtcdV3Backend
type jsiiProxy_EtcdV3Backend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_EtcdV3Backend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EtcdV3Backend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewEtcdV3Backend(scope constructs.Construct, props *EtcdV3BackendProps) EtcdV3Backend {
	_init_.Initialize()

	j := jsiiProxy_EtcdV3Backend{}

	_jsii_.Create(
		"cdktf.EtcdV3Backend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEtcdV3Backend_Override(e EtcdV3Backend, scope constructs.Construct, props *EtcdV3BackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.EtcdV3Backend",
		[]interface{}{scope, props},
		e,
	)
}

// Experimental.
func (e *jsiiProxy_EtcdV3Backend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EtcdV3Backend) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EtcdV3Backend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (e *jsiiProxy_EtcdV3Backend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EtcdV3Backend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (e *jsiiProxy_EtcdV3Backend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EtcdV3Backend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_EtcdV3Backend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type EtcdV3BackendProps struct {
	// Experimental.
	Endpoints *[]*string `json:"endpoints"`
	// Experimental.
	CacertPath *string `json:"cacertPath"`
	// Experimental.
	CertPath *string `json:"certPath"`
	// Experimental.
	KeyPath *string `json:"keyPath"`
	// Experimental.
	Lock *bool `json:"lock"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Username *string `json:"username"`
}

// Experimental.
type GcsBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for GcsBackend
type jsiiProxy_GcsBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_GcsBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GcsBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewGcsBackend(scope constructs.Construct, props *GcsBackendProps) GcsBackend {
	_init_.Initialize()

	j := jsiiProxy_GcsBackend{}

	_jsii_.Create(
		"cdktf.GcsBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGcsBackend_Override(g GcsBackend, scope constructs.Construct, props *GcsBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.GcsBackend",
		[]interface{}{scope, props},
		g,
	)
}

// Experimental.
func (g *jsiiProxy_GcsBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (g *jsiiProxy_GcsBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		g,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GcsBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (g *jsiiProxy_GcsBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (g *jsiiProxy_GcsBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (g *jsiiProxy_GcsBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (g *jsiiProxy_GcsBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (g *jsiiProxy_GcsBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type GcsBackendProps struct {
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	AccessToken *string `json:"accessToken"`
	// Experimental.
	Credentials *string `json:"credentials"`
	// Experimental.
	EncryptionKey *string `json:"encryptionKey"`
	// Experimental.
	Prefix *string `json:"prefix"`
}

// Experimental.
type HttpBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for HttpBackend
type jsiiProxy_HttpBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_HttpBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpBackend(scope constructs.Construct, props *HttpBackendProps) HttpBackend {
	_init_.Initialize()

	j := jsiiProxy_HttpBackend{}

	_jsii_.Create(
		"cdktf.HttpBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpBackend_Override(h HttpBackend, scope constructs.Construct, props *HttpBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.HttpBackend",
		[]interface{}{scope, props},
		h,
	)
}

// Experimental.
func (h *jsiiProxy_HttpBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (h *jsiiProxy_HttpBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (h *jsiiProxy_HttpBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (h *jsiiProxy_HttpBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (h *jsiiProxy_HttpBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (h *jsiiProxy_HttpBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (h *jsiiProxy_HttpBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (h *jsiiProxy_HttpBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type HttpBackendProps struct {
	// Experimental.
	Address *string `json:"address"`
	// Experimental.
	LockAddress *string `json:"lockAddress"`
	// Experimental.
	LockMethod *string `json:"lockMethod"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	RetryMax *float64 `json:"retryMax"`
	// Experimental.
	RetryWaitMax *float64 `json:"retryWaitMax"`
	// Experimental.
	RetryWaitMin *float64 `json:"retryWaitMin"`
	// Experimental.
	SkipCertVerification *bool `json:"skipCertVerification"`
	// Experimental.
	UnlockAddress *string `json:"unlockAddress"`
	// Experimental.
	UnlockMethod *string `json:"unlockMethod"`
	// Experimental.
	UpdateMethod *string `json:"updateMethod"`
	// Experimental.
	Username *string `json:"username"`
}

// Interface for lazy untyped value producers.
// Experimental.
type IAnyProducer interface {
	// Produce the value.
	// Experimental.
	Produce(context IResolveContext) interface{}
}

// The jsii proxy for IAnyProducer
type jsiiProxy_IAnyProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IAnyProducer) Produce(context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// Function used to concatenate symbols in the target document language.
//
// Interface so it could potentially be exposed over jsii.
// Experimental.
type IFragmentConcatenator interface {
	// Join the fragment on the left and on the right.
	// Experimental.
	Join(left interface{}, right interface{}) interface{}
}

// The jsii proxy for IFragmentConcatenator
type jsiiProxy_IFragmentConcatenator struct {
	_ byte // padding
}

func (i *jsiiProxy_IFragmentConcatenator) Join(left interface{}, right interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"join",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Interface for lazy list producers.
// Experimental.
type IListProducer interface {
	// Produce the list value.
	// Experimental.
	Produce(context IResolveContext) *[]*string
}

// The jsii proxy for IListProducer
type jsiiProxy_IListProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IListProducer) Produce(context IResolveContext) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// Interface for lazy number producers.
// Experimental.
type INumberProducer interface {
	// Produce the number value.
	// Experimental.
	Produce(context IResolveContext) *float64
}

// The jsii proxy for INumberProducer
type jsiiProxy_INumberProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_INumberProducer) Produce(context IResolveContext) *float64 {
	var returns *float64

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// A Token that can post-process the complete resolved value, after resolve() has recursed over it.
// Experimental.
type IPostProcessor interface {
	// Process the completely resolved value, after full recursion/resolution has happened.
	// Experimental.
	PostProcess(input interface{}, context IResolveContext) interface{}
}

// The jsii proxy for IPostProcessor
type jsiiProxy_IPostProcessor struct {
	_ byte // padding
}

func (i *jsiiProxy_IPostProcessor) PostProcess(input interface{}, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"postProcess",
		[]interface{}{input, context},
		&returns,
	)

	return returns
}

// Experimental.
type IRemoteWorkspace interface {
}

// The jsii proxy for IRemoteWorkspace
type jsiiProxy_IRemoteWorkspace struct {
	_ byte // padding
}

// Interface for values that can be resolvable later.
//
// Tokens are special objects that participate in synthesis.
// Experimental.
type IResolvable interface {
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
}

// The jsii proxy for IResolvable
type jsiiProxy_IResolvable struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolvable) Resolve(context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IResolvable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResolvable) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

// Current resolution context for tokens.
// Experimental.
type IResolveContext interface {
	// Use this postprocessor after the entire token structure has been resolved.
	// Experimental.
	RegisterPostProcessor(postProcessor IPostProcessor)
	// Resolve an inner object.
	// Experimental.
	Resolve(x interface{}) interface{}
	// True when we are still preparing, false if we're rendering the final output.
	// Experimental.
	Preparing() *bool
	// The scope from which resolution has been initiated.
	// Experimental.
	Scope() constructs.IConstruct
}

// The jsii proxy for IResolveContext
type jsiiProxy_IResolveContext struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolveContext) RegisterPostProcessor(postProcessor IPostProcessor) {
	_jsii_.InvokeVoid(
		i,
		"registerPostProcessor",
		[]interface{}{postProcessor},
	)
}

func (i *jsiiProxy_IResolveContext) Resolve(x interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResolveContext) Preparing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"preparing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolveContext) Scope() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

// Experimental.
type IResource interface {
	constructs.IConstruct
	// The stack in which this resource is defined.
	// Experimental.
	Stack() TerraformStack
}

// The jsii proxy for IResource
type jsiiProxy_IResource struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResource) Stack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Interface for lazy string producers.
// Experimental.
type IStringProducer interface {
	// Produce the string value.
	// Experimental.
	Produce(context IResolveContext) *string
}

// The jsii proxy for IStringProducer
type jsiiProxy_IStringProducer struct {
	_ byte // padding
}

func (i *jsiiProxy_IStringProducer) Produce(context IResolveContext) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"produce",
		[]interface{}{context},
		&returns,
	)

	return returns
}

// Experimental.
type ITerraformDependable interface {
	// Experimental.
	Fqn() *string
}

// The jsii proxy for ITerraformDependable
type jsiiProxy_ITerraformDependable struct {
	_ byte // padding
}

func (j *jsiiProxy_ITerraformDependable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

// Experimental.
type ITerraformResource interface {
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) *string
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(c *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(d *[]*string)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(l *TerraformResourceLifecycle)
	// Experimental.
	Provider() TerraformProvider
	// Experimental.
	SetProvider(p TerraformProvider)
	// Experimental.
	TerraformResourceType() *string
}

// The jsii proxy for ITerraformResource
type jsiiProxy_ITerraformResource struct {
	_ byte // padding
}

func (i *jsiiProxy_ITerraformResource) InterpolationForAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITerraformResource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ITerraformResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ITerraformResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) Lifecycle() *TerraformResourceLifecycle {
	var returns *TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) SetLifecycle(val *TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ITerraformResource) Provider() TerraformProvider {
	var returns TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITerraformResource) SetProvider(val TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ITerraformResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Interface to apply operation to tokens in a string.
//
// Interface so it can be exported via jsii.
// Experimental.
type ITokenMapper interface {
	// Replace a single token.
	// Experimental.
	MapToken(t IResolvable) interface{}
}

// The jsii proxy for ITokenMapper
type jsiiProxy_ITokenMapper struct {
	_ byte // padding
}

func (i *jsiiProxy_ITokenMapper) MapToken(t IResolvable) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"mapToken",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// How to resolve tokens.
// Experimental.
type ITokenResolver interface {
	// Resolve a tokenized list.
	// Experimental.
	ResolveList(l *[]*string, context IResolveContext) interface{}
	// Resolve a string with at least one stringified token in it.
	//
	// (May use concatenation)
	// Experimental.
	ResolveString(s TokenizedStringFragments, context IResolveContext) interface{}
	// Resolve a single token.
	// Experimental.
	ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{}
}

// The jsii proxy for ITokenResolver
type jsiiProxy_ITokenResolver struct {
	_ byte // padding
}

func (i *jsiiProxy_ITokenResolver) ResolveList(l *[]*string, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveList",
		[]interface{}{l, context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITokenResolver) ResolveString(s TokenizedStringFragments, context IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveString",
		[]interface{}{s, context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITokenResolver) ResolveToken(t IResolvable, context IResolveContext, postProcessor IPostProcessor) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolveToken",
		[]interface{}{t, context, postProcessor},
		&returns,
	)

	return returns
}

// Lazily produce a value.
//
// Can be used to return a string, list or numeric value whose actual value
// will only be calculated later, during synthesis.
// Experimental.
type Lazy interface {
}

// The jsii proxy struct for Lazy
type jsiiProxy_Lazy struct {
	_ byte // padding
}

// Experimental.
func NewLazy() Lazy {
	_init_.Initialize()

	j := jsiiProxy_Lazy{}

	_jsii_.Create(
		"cdktf.Lazy",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewLazy_Override(l Lazy) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Lazy",
		nil, // no parameters
		l,
	)
}

// Produces a lazy token from an untyped value.
// Experimental.
func Lazy_AnyValue(producer IAnyProducer, options *LazyAnyValueOptions) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Lazy",
		"anyValue",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Returns a list-ified token for a lazy value.
// Experimental.
func Lazy_ListValue(producer IListProducer, options *LazyListValueOptions) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Lazy",
		"listValue",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Returns a numberified token for a lazy value.
// Experimental.
func Lazy_NumberValue(producer INumberProducer) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Lazy",
		"numberValue",
		[]interface{}{producer},
		&returns,
	)

	return returns
}

// Returns a stringified token for a lazy value.
// Experimental.
func Lazy_StringValue(producer IStringProducer, options *LazyStringValueOptions) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Lazy",
		"stringValue",
		[]interface{}{producer, options},
		&returns,
	)

	return returns
}

// Options for creating lazy untyped tokens.
// Experimental.
type LazyAnyValueOptions struct {
	// Use the given name as a display hint.
	// Experimental.
	DisplayHint *string `json:"displayHint"`
	// If the produced value is an array and it is empty, return 'undefined' instead.
	// Experimental.
	OmitEmptyArray *bool `json:"omitEmptyArray"`
}

// Options for creating a lazy list token.
// Experimental.
type LazyListValueOptions struct {
	// Use the given name as a display hint.
	// Experimental.
	DisplayHint *string `json:"displayHint"`
	// If the produced list is empty, return 'undefined' instead.
	// Experimental.
	OmitEmpty *bool `json:"omitEmpty"`
}

// Options for creating a lazy string token.
// Experimental.
type LazyStringValueOptions struct {
	// Use the given name as a display hint.
	// Experimental.
	DisplayHint *string `json:"displayHint"`
}

// Experimental.
type LocalBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LocalBackend
type jsiiProxy_LocalBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_LocalBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LocalBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewLocalBackend(scope constructs.Construct, props *LocalBackendProps) LocalBackend {
	_init_.Initialize()

	j := jsiiProxy_LocalBackend{}

	_jsii_.Create(
		"cdktf.LocalBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLocalBackend_Override(l LocalBackend, scope constructs.Construct, props *LocalBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.LocalBackend",
		[]interface{}{scope, props},
		l,
	)
}

// Experimental.
func (l *jsiiProxy_LocalBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_LocalBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LocalBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (l *jsiiProxy_LocalBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (l *jsiiProxy_LocalBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (l *jsiiProxy_LocalBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (l *jsiiProxy_LocalBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (l *jsiiProxy_LocalBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type LocalBackendProps struct {
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	WorkspaceDir *string `json:"workspaceDir"`
}

// Experimental.
type Manifest interface {
	Outdir() *string
	Stacks() *[]*StackManifest
	Version() *string
	BuildManifest() interface{}
	ForStack(stack TerraformStack) *StackManifest
	WriteToFile()
}

// The jsii proxy struct for Manifest
type jsiiProxy_Manifest struct {
	_ byte // padding
}

func (j *jsiiProxy_Manifest) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Manifest) Stacks() *[]*StackManifest {
	var returns *[]*StackManifest
	_jsii_.Get(
		j,
		"stacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Manifest) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewManifest(version *string, outdir *string) Manifest {
	_init_.Initialize()

	j := jsiiProxy_Manifest{}

	_jsii_.Create(
		"cdktf.Manifest",
		[]interface{}{version, outdir},
		&j,
	)

	return &j
}

// Experimental.
func NewManifest_Override(m Manifest, version *string, outdir *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Manifest",
		[]interface{}{version, outdir},
		m,
	)
}

func Manifest_FileName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.Manifest",
		"fileName",
		&returns,
	)
	return returns
}

func Manifest_StackFileName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.Manifest",
		"stackFileName",
		&returns,
	)
	return returns
}

func Manifest_StacksFolder() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.Manifest",
		"stacksFolder",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_Manifest) BuildManifest() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"buildManifest",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Manifest) ForStack(stack TerraformStack) *StackManifest {
	var returns *StackManifest

	_jsii_.Invoke(
		m,
		"forStack",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_Manifest) WriteToFile() {
	_jsii_.InvokeVoid(
		m,
		"writeToFile",
		nil, // no parameters
	)
}

// Experimental.
type MantaBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for MantaBackend
type jsiiProxy_MantaBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_MantaBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MantaBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewMantaBackend(scope constructs.Construct, props *MantaBackendProps) MantaBackend {
	_init_.Initialize()

	j := jsiiProxy_MantaBackend{}

	_jsii_.Create(
		"cdktf.MantaBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMantaBackend_Override(m MantaBackend, scope constructs.Construct, props *MantaBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.MantaBackend",
		[]interface{}{scope, props},
		m,
	)
}

// Experimental.
func (m *jsiiProxy_MantaBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (m *jsiiProxy_MantaBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		m,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (m *jsiiProxy_MantaBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		m,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (m *jsiiProxy_MantaBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_MantaBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (m *jsiiProxy_MantaBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (m *jsiiProxy_MantaBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (m *jsiiProxy_MantaBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type MantaBackendProps struct {
	// Experimental.
	Account *string `json:"account"`
	// Experimental.
	KeyId *string `json:"keyId"`
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	InsecureSkipTlsVerify *bool `json:"insecureSkipTlsVerify"`
	// Experimental.
	KeyMaterial *string `json:"keyMaterial"`
	// Experimental.
	ObjectName *string `json:"objectName"`
	// Experimental.
	Url *string `json:"url"`
	// Experimental.
	User *string `json:"user"`
}

// Experimental.
type NamedRemoteWorkspace interface {
	IRemoteWorkspace
	Name() *string
	SetName(val *string)
}

// The jsii proxy struct for NamedRemoteWorkspace
type jsiiProxy_NamedRemoteWorkspace struct {
	jsiiProxy_IRemoteWorkspace
}

func (j *jsiiProxy_NamedRemoteWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewNamedRemoteWorkspace(name *string) NamedRemoteWorkspace {
	_init_.Initialize()

	j := jsiiProxy_NamedRemoteWorkspace{}

	_jsii_.Create(
		"cdktf.NamedRemoteWorkspace",
		[]interface{}{name},
		&j,
	)

	return &j
}

// Experimental.
func NewNamedRemoteWorkspace_Override(n NamedRemoteWorkspace, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.NamedRemoteWorkspace",
		[]interface{}{name},
		n,
	)
}

func (j *jsiiProxy_NamedRemoteWorkspace) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Experimental.
type NumberMap interface {
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() ITerraformResource
	SetTerraformResource(val ITerraformResource)
	Lookup(key *string) *float64
}

// The jsii proxy struct for NumberMap
type jsiiProxy_NumberMap struct {
	_ byte // padding
}

func (j *jsiiProxy_NumberMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NumberMap) TerraformResource() ITerraformResource {
	var returns ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewNumberMap(terraformResource ITerraformResource, terraformAttribute *string) NumberMap {
	_init_.Initialize()

	j := jsiiProxy_NumberMap{}

	_jsii_.Create(
		"cdktf.NumberMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewNumberMap_Override(n NumberMap, terraformResource ITerraformResource, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.NumberMap",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NumberMap) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NumberMap) SetTerraformResource(val ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (n *jsiiProxy_NumberMap) Lookup(key *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"lookup",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Experimental.
type OssAssumeRole struct {
	// Experimental.
	RoleArn *string `json:"roleArn"`
	// Experimental.
	Policy *string `json:"policy"`
	// Experimental.
	SessionExpiration *float64 `json:"sessionExpiration"`
	// Experimental.
	SessionName *string `json:"sessionName"`
}

// Experimental.
type OssBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OssBackend
type jsiiProxy_OssBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_OssBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OssBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewOssBackend(scope constructs.Construct, props *OssBackendProps) OssBackend {
	_init_.Initialize()

	j := jsiiProxy_OssBackend{}

	_jsii_.Create(
		"cdktf.OssBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewOssBackend_Override(o OssBackend, scope constructs.Construct, props *OssBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.OssBackend",
		[]interface{}{scope, props},
		o,
	)
}

// Experimental.
func (o *jsiiProxy_OssBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (o *jsiiProxy_OssBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		o,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (o *jsiiProxy_OssBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		o,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (o *jsiiProxy_OssBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OssBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (o *jsiiProxy_OssBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (o *jsiiProxy_OssBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OssBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type OssBackendProps struct {
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	AssumeRole *OssAssumeRole `json:"assumeRole"`
	// Experimental.
	EcsRoleName *string `json:"ecsRoleName"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Prefix *string `json:"prefix"`
	// Experimental.
	Profile *string `json:"profile"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
	// Experimental.
	SecurityToken *string `json:"securityToken"`
	// Experimental.
	SharedCredentialsFile *string `json:"sharedCredentialsFile"`
	// Experimental.
	TablestoreEndpoint *string `json:"tablestoreEndpoint"`
	// Experimental.
	TablestoreTable *string `json:"tablestoreTable"`
}

// Experimental.
type PgBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for PgBackend
type jsiiProxy_PgBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_PgBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PgBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewPgBackend(scope constructs.Construct, props *PgBackendProps) PgBackend {
	_init_.Initialize()

	j := jsiiProxy_PgBackend{}

	_jsii_.Create(
		"cdktf.PgBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPgBackend_Override(p PgBackend, scope constructs.Construct, props *PgBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.PgBackend",
		[]interface{}{scope, props},
		p,
	)
}

// Experimental.
func (p *jsiiProxy_PgBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (p *jsiiProxy_PgBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (p *jsiiProxy_PgBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (p *jsiiProxy_PgBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (p *jsiiProxy_PgBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (p *jsiiProxy_PgBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_PgBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (p *jsiiProxy_PgBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type PgBackendProps struct {
	// Experimental.
	ConnStr *string `json:"connStr"`
	// Experimental.
	SchemaName *string `json:"schemaName"`
	// Experimental.
	SkipSchemaCreation *bool `json:"skipSchemaCreation"`
}

// Experimental.
type PrefixedRemoteWorkspaces interface {
	IRemoteWorkspace
	Prefix() *string
	SetPrefix(val *string)
}

// The jsii proxy struct for PrefixedRemoteWorkspaces
type jsiiProxy_PrefixedRemoteWorkspaces struct {
	jsiiProxy_IRemoteWorkspace
}

func (j *jsiiProxy_PrefixedRemoteWorkspaces) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}


// Experimental.
func NewPrefixedRemoteWorkspaces(prefix *string) PrefixedRemoteWorkspaces {
	_init_.Initialize()

	j := jsiiProxy_PrefixedRemoteWorkspaces{}

	_jsii_.Create(
		"cdktf.PrefixedRemoteWorkspaces",
		[]interface{}{prefix},
		&j,
	)

	return &j
}

// Experimental.
func NewPrefixedRemoteWorkspaces_Override(p PrefixedRemoteWorkspaces, prefix *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.PrefixedRemoteWorkspaces",
		[]interface{}{prefix},
		p,
	)
}

func (j *jsiiProxy_PrefixedRemoteWorkspaces) SetPrefix(val *string) {
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

// Experimental.
type RemoteBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for RemoteBackend
type jsiiProxy_RemoteBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_RemoteBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoteBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewRemoteBackend(scope constructs.Construct, props *RemoteBackendProps) RemoteBackend {
	_init_.Initialize()

	j := jsiiProxy_RemoteBackend{}

	_jsii_.Create(
		"cdktf.RemoteBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRemoteBackend_Override(r RemoteBackend, scope constructs.Construct, props *RemoteBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.RemoteBackend",
		[]interface{}{scope, props},
		r,
	)
}

// Experimental.
func (r *jsiiProxy_RemoteBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (r *jsiiProxy_RemoteBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (r *jsiiProxy_RemoteBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (r *jsiiProxy_RemoteBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (r *jsiiProxy_RemoteBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (r *jsiiProxy_RemoteBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (r *jsiiProxy_RemoteBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (r *jsiiProxy_RemoteBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type RemoteBackendProps struct {
	// Experimental.
	Organization *string `json:"organization"`
	// Experimental.
	Workspaces IRemoteWorkspace `json:"workspaces"`
	// Experimental.
	Hostname *string `json:"hostname"`
	// Experimental.
	Token *string `json:"token"`
}

// Options to the resolve() operation.
//
// NOT the same as the ResolveContext; ResolveContext is exposed to Token
// implementors and resolution hooks, whereas this struct is just to bundle
// a number of things that would otherwise be arguments to resolve() in a
// readable way.
// Experimental.
type ResolveOptions struct {
	// The resolver to apply to any resolvable tokens found.
	// Experimental.
	Resolver ITokenResolver `json:"resolver"`
	// The scope from which resolution is performed.
	// Experimental.
	Scope constructs.IConstruct `json:"scope"`
	// Whether the resolution is being executed during the prepare phase or not.
	// Experimental.
	Preparing *bool `json:"preparing"`
}

// A construct which represents a resource.
// Experimental.
type Resource interface {
	constructs.Construct
	IResource
	Stack() TerraformStack
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToString() *string
}

// The jsii proxy struct for Resource
type jsiiProxy_Resource struct {
	internal.Type__constructsConstruct
	jsiiProxy_IResource
}

func (j *jsiiProxy_Resource) Stack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewResource_Override(r Resource, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Resource",
		[]interface{}{scope, id},
		r,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (r *jsiiProxy_Resource) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (r *jsiiProxy_Resource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (r *jsiiProxy_Resource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (r *jsiiProxy_Resource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type S3Backend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for S3Backend
type jsiiProxy_S3Backend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_S3Backend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3Backend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewS3Backend(scope constructs.Construct, props *S3BackendProps) S3Backend {
	_init_.Initialize()

	j := jsiiProxy_S3Backend{}

	_jsii_.Create(
		"cdktf.S3Backend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Backend_Override(s S3Backend, scope constructs.Construct, props *S3BackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.S3Backend",
		[]interface{}{scope, props},
		s,
	)
}

// Experimental.
func (s *jsiiProxy_S3Backend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_S3Backend) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_S3Backend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (s *jsiiProxy_S3Backend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_S3Backend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (s *jsiiProxy_S3Backend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_S3Backend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_S3Backend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type S3BackendProps struct {
	// Experimental.
	Bucket *string `json:"bucket"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	AccessKey *string `json:"accessKey"`
	// Experimental.
	Acl *string `json:"acl"`
	// Experimental.
	AssumeRolePolicy *string `json:"assumeRolePolicy"`
	// Experimental.
	DynamodbEndpoint *string `json:"dynamodbEndpoint"`
	// Experimental.
	DynamodbTable *string `json:"dynamodbTable"`
	// Experimental.
	Encrypt *bool `json:"encrypt"`
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// Experimental.
	ExternalId *string `json:"externalId"`
	// Experimental.
	ForcePathStyle *bool `json:"forcePathStyle"`
	// Experimental.
	IamEndpoint *string `json:"iamEndpoint"`
	// Experimental.
	KmsKeyId *string `json:"kmsKeyId"`
	// Experimental.
	MaxRetries *float64 `json:"maxRetries"`
	// Experimental.
	Profile *string `json:"profile"`
	// Experimental.
	Region *string `json:"region"`
	// Experimental.
	RoleArn *string `json:"roleArn"`
	// Experimental.
	SecretKey *string `json:"secretKey"`
	// Experimental.
	SessionName *string `json:"sessionName"`
	// Experimental.
	SharedCredentialsFile *string `json:"sharedCredentialsFile"`
	// Experimental.
	SkipCredentialsValidation *bool `json:"skipCredentialsValidation"`
	// Experimental.
	SkipMetadataApiCheck *bool `json:"skipMetadataApiCheck"`
	// Experimental.
	SseCustomerKey *string `json:"sseCustomerKey"`
	// Experimental.
	StsEndpoint *string `json:"stsEndpoint"`
	// Experimental.
	Token *string `json:"token"`
	// Experimental.
	WorkspaceKeyPrefix *string `json:"workspaceKeyPrefix"`
}

// Experimental.
type StackManifest struct {
	// Experimental.
	ConstructPath *string `json:"constructPath"`
	// Experimental.
	Name *string `json:"name"`
	// Experimental.
	SynthesizedStackPath *string `json:"synthesizedStackPath"`
	// Experimental.
	WorkingDirectory *string `json:"workingDirectory"`
}

// Converts all fragments to strings and concats those.
//
// Drops 'undefined's.
// Experimental.
type StringConcat interface {
	IFragmentConcatenator
	Join(left interface{}, right interface{}) interface{}
}

// The jsii proxy struct for StringConcat
type jsiiProxy_StringConcat struct {
	jsiiProxy_IFragmentConcatenator
}

// Experimental.
func NewStringConcat() StringConcat {
	_init_.Initialize()

	j := jsiiProxy_StringConcat{}

	_jsii_.Create(
		"cdktf.StringConcat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStringConcat_Override(s StringConcat) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.StringConcat",
		nil, // no parameters
		s,
	)
}

// Join the fragment on the left and on the right.
// Experimental.
func (s *jsiiProxy_StringConcat) Join(left interface{}, right interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"join",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Experimental.
type StringMap interface {
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() ITerraformResource
	SetTerraformResource(val ITerraformResource)
	Lookup(key *string) *string
}

// The jsii proxy struct for StringMap
type jsiiProxy_StringMap struct {
	_ byte // padding
}

func (j *jsiiProxy_StringMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StringMap) TerraformResource() ITerraformResource {
	var returns ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewStringMap(terraformResource ITerraformResource, terraformAttribute *string) StringMap {
	_init_.Initialize()

	j := jsiiProxy_StringMap{}

	_jsii_.Create(
		"cdktf.StringMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewStringMap_Override(s StringMap, terraformResource ITerraformResource, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.StringMap",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StringMap) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StringMap) SetTerraformResource(val ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_StringMap) Lookup(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"lookup",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Experimental.
type SwiftBackend interface {
	TerraformBackend
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SwiftBackend
type jsiiProxy_SwiftBackend struct {
	jsiiProxy_TerraformBackend
}

func (j *jsiiProxy_SwiftBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SwiftBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewSwiftBackend(scope constructs.Construct, props *SwiftBackendProps) SwiftBackend {
	_init_.Initialize()

	j := jsiiProxy_SwiftBackend{}

	_jsii_.Create(
		"cdktf.SwiftBackend",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSwiftBackend_Override(s SwiftBackend, scope constructs.Construct, props *SwiftBackendProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.SwiftBackend",
		[]interface{}{scope, props},
		s,
	)
}

// Experimental.
func (s *jsiiProxy_SwiftBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SwiftBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SwiftBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (s *jsiiProxy_SwiftBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_SwiftBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (s *jsiiProxy_SwiftBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_SwiftBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_SwiftBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type SwiftBackendProps struct {
	// Experimental.
	Container *string `json:"container"`
	// Experimental.
	ApplicationCredentialId *string `json:"applicationCredentialId"`
	// Experimental.
	ApplicationCredentialName *string `json:"applicationCredentialName"`
	// Experimental.
	ApplicationCredentialSecret *string `json:"applicationCredentialSecret"`
	// Experimental.
	ArchiveContainer *string `json:"archiveContainer"`
	// Experimental.
	AuthUrl *string `json:"authUrl"`
	// Experimental.
	CacertFile *string `json:"cacertFile"`
	// Experimental.
	Cert *string `json:"cert"`
	// Experimental.
	Cloud *string `json:"cloud"`
	// Experimental.
	DefaultDomain *string `json:"defaultDomain"`
	// Experimental.
	DomainId *string `json:"domainId"`
	// Experimental.
	DomainName *string `json:"domainName"`
	// Experimental.
	ExpireAfter *string `json:"expireAfter"`
	// Experimental.
	Insecure *bool `json:"insecure"`
	// Experimental.
	Key *string `json:"key"`
	// Experimental.
	Password *string `json:"password"`
	// Experimental.
	ProjectDomainId *string `json:"projectDomainId"`
	// Experimental.
	ProjectDomainName *string `json:"projectDomainName"`
	// Experimental.
	RegionName *string `json:"regionName"`
	// Experimental.
	StateName *string `json:"stateName"`
	// Experimental.
	TenantId *string `json:"tenantId"`
	// Experimental.
	TenantName *string `json:"tenantName"`
	// Experimental.
	Token *string `json:"token"`
	// Experimental.
	UserDomainId *string `json:"userDomainId"`
	// Experimental.
	UserDomainName *string `json:"userDomainName"`
	// Experimental.
	UserId *string `json:"userId"`
	// Experimental.
	UserName *string `json:"userName"`
}

// Experimental.
type TerraformAsset interface {
	Resource
	AssetHash() *string
	SetAssetHash(val *string)
	FileName() *string
	Path() *string
	Stack() TerraformStack
	Type() AssetType
	SetType(val AssetType)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToString() *string
}

// The jsii proxy struct for TerraformAsset
type jsiiProxy_TerraformAsset struct {
	jsiiProxy_Resource
}

func (j *jsiiProxy_TerraformAsset) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformAsset) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformAsset) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformAsset) Stack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformAsset) Type() AssetType {
	var returns AssetType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// A Terraform Asset takes a file or directory outside of the CDK for Terraform context and moves it into it.
//
// Assets copy referenced files into the stacks context for further usage in other resources.
// Experimental.
func NewTerraformAsset(scope constructs.Construct, id *string, config *TerraformAssetConfig) TerraformAsset {
	_init_.Initialize()

	j := jsiiProxy_TerraformAsset{}

	_jsii_.Create(
		"cdktf.TerraformAsset",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// A Terraform Asset takes a file or directory outside of the CDK for Terraform context and moves it into it.
//
// Assets copy referenced files into the stacks context for further usage in other resources.
// Experimental.
func NewTerraformAsset_Override(t TerraformAsset, scope constructs.Construct, id *string, config *TerraformAssetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformAsset",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TerraformAsset) SetAssetHash(val *string) {
	_jsii_.Set(
		j,
		"assetHash",
		val,
	)
}

func (j *jsiiProxy_TerraformAsset) SetType(val AssetType) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformAsset) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformAsset) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformAsset) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformAssetConfig struct {
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	AssetHash *string `json:"assetHash"`
	// Experimental.
	Type AssetType `json:"type"`
}

// Experimental.
type TerraformBackend interface {
	TerraformElement
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	Name() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformBackend
type jsiiProxy_TerraformBackend struct {
	jsiiProxy_TerraformElement
}

func (j *jsiiProxy_TerraformBackend) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformBackend_Override(t TerraformBackend, scope constructs.Construct, id *string, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformBackend",
		[]interface{}{scope, id, name},
		t,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformBackend) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformBackend) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformBackend) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformBackend) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformBackend) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (t *jsiiProxy_TerraformBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformDataSource interface {
	TerraformElement
	ITerraformDependable
	ITerraformResource
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Lifecycle() *TerraformResourceLifecycle
	SetLifecycle(val *TerraformResourceLifecycle)
	Provider() TerraformProvider
	SetProvider(val TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) *bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformDataSource
type jsiiProxy_TerraformDataSource struct {
	jsiiProxy_TerraformElement
	jsiiProxy_ITerraformDependable
	jsiiProxy_ITerraformResource
}

func (j *jsiiProxy_TerraformDataSource) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) Lifecycle() *TerraformResourceLifecycle {
	var returns *TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) Provider() TerraformProvider {
	var returns TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata {
	var returns *TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformDataSource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformDataSource(scope constructs.Construct, id *string, config *TerraformResourceConfig) TerraformDataSource {
	_init_.Initialize()

	j := jsiiProxy_TerraformDataSource{}

	_jsii_.Create(
		"cdktf.TerraformDataSource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformDataSource_Override(t TerraformDataSource, scope constructs.Construct, id *string, config *TerraformResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformDataSource",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TerraformDataSource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TerraformDataSource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TerraformDataSource) SetLifecycle(val *TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TerraformDataSource) SetProvider(val TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformDataSource) GetBooleanAttribute(terraformAttribute *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformDataSource) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformDataSource) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformDataSource) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformDataSource) InterpolationForAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformDataSource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (t *jsiiProxy_TerraformDataSource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformElement interface {
	constructs.Construct
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformElement
type jsiiProxy_TerraformElement struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_TerraformElement) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformElement) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformElement) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformElement) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformElement) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformElement(scope constructs.Construct, id *string) TerraformElement {
	_init_.Initialize()

	j := jsiiProxy_TerraformElement{}

	_jsii_.Create(
		"cdktf.TerraformElement",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformElement_Override(t TerraformElement, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformElement",
		[]interface{}{scope, id},
		t,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformElement) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformElement) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformElement) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformElement) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformElement) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformElement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformElement) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformElementMetadata struct {
	// Experimental.
	Path *string `json:"path"`
	// Experimental.
	StackTrace *[]*string `json:"stackTrace"`
	// Experimental.
	UniqueId *string `json:"uniqueId"`
}

// Experimental.
type TerraformHclModule interface {
	TerraformModule
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Providers() *[]interface{}
	RawOverrides() interface{}
	Source() *string
	Variables() *map[string]interface{}
	Version() *string
	AddOverride(path *string, value interface{})
	AddProvider(provider interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	InterpolationForOutput(moduleOutput *string) interface{}
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Set(variable *string, value interface{})
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformHclModule
type jsiiProxy_TerraformHclModule struct {
	jsiiProxy_TerraformModule
}

func (j *jsiiProxy_TerraformHclModule) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Variables() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformHclModule) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformHclModule(scope constructs.Construct, id *string, options *TerraformHclModuleOptions) TerraformHclModule {
	_init_.Initialize()

	j := jsiiProxy_TerraformHclModule{}

	_jsii_.Create(
		"cdktf.TerraformHclModule",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformHclModule_Override(t TerraformHclModule, scope constructs.Construct, id *string, options *TerraformHclModuleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformHclModule",
		[]interface{}{scope, id, options},
		t,
	)
}

func (j *jsiiProxy_TerraformHclModule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) AddProvider(provider interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addProvider",
		[]interface{}{provider},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) InterpolationForOutput(moduleOutput *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformHclModule) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformHclModule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformHclModule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformHclModule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) Set(variable *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"set",
		[]interface{}{variable, value},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformHclModule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformHclModule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformHclModuleOptions struct {
	// Experimental.
	Source *string `json:"source"`
	// Experimental.
	DependsOn *[]ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Providers *[]interface{} `json:"providers"`
	// Experimental.
	Version *string `json:"version"`
	// Experimental.
	Variables *map[string]interface{} `json:"variables"`
}

// Experimental.
type TerraformLocal interface {
	TerraformElement
	AsBoolean() *bool
	AsList() *[]*string
	AsNumber() *float64
	AsString() *string
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	Expression() interface{}
	SetExpression(val interface{})
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformLocal
type jsiiProxy_TerraformLocal struct {
	jsiiProxy_TerraformElement
}

func (j *jsiiProxy_TerraformLocal) AsBoolean() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"asBoolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) AsList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) AsNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) AsString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) Expression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformLocal) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformLocal(scope constructs.Construct, id *string, expression interface{}) TerraformLocal {
	_init_.Initialize()

	j := jsiiProxy_TerraformLocal{}

	_jsii_.Create(
		"cdktf.TerraformLocal",
		[]interface{}{scope, id, expression},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformLocal_Override(t TerraformLocal, scope constructs.Construct, id *string, expression interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformLocal",
		[]interface{}{scope, id, expression},
		t,
	)
}

func (j *jsiiProxy_TerraformLocal) SetExpression(val interface{}) {
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformLocal) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformLocal) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformLocal) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformLocal) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformLocal) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformLocal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformLocal) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformMetaArguments struct {
	// Experimental.
	Count *float64 `json:"count"`
	// Experimental.
	DependsOn *[]ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider TerraformProvider `json:"provider"`
}

// Experimental.
type TerraformModule interface {
	TerraformElement
	ITerraformDependable
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Providers() *[]interface{}
	RawOverrides() interface{}
	Source() *string
	Version() *string
	AddOverride(path *string, value interface{})
	AddProvider(provider interface{})
	InterpolationForOutput(moduleOutput *string) interface{}
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformModule
type jsiiProxy_TerraformModule struct {
	jsiiProxy_TerraformElement
	jsiiProxy_ITerraformDependable
}

func (j *jsiiProxy_TerraformModule) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformModule) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformModule_Override(t TerraformModule, scope constructs.Construct, id *string, options *TerraformModuleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformModule",
		[]interface{}{scope, id, options},
		t,
	)
}

func (j *jsiiProxy_TerraformModule) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformModule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformModule) AddProvider(provider interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addProvider",
		[]interface{}{provider},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformModule) InterpolationForOutput(moduleOutput *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformModule) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformModule) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformModule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformModule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformModule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformModule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformModule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformModuleOptions struct {
	// Experimental.
	Source *string `json:"source"`
	// Experimental.
	DependsOn *[]ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Providers *[]interface{} `json:"providers"`
	// Experimental.
	Version *string `json:"version"`
}

// Experimental.
type TerraformModuleProvider struct {
	// Experimental.
	ModuleAlias *string `json:"moduleAlias"`
	// Experimental.
	Provider TerraformProvider `json:"provider"`
}

// Experimental.
type TerraformOutput interface {
	TerraformElement
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	DependsOn() *[]ITerraformDependable
	SetDependsOn(val *[]ITerraformDependable)
	Description() *string
	SetDescription(val *string)
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	Sensitive() *bool
	SetSensitive(val *bool)
	Value() interface{}
	SetValue(val interface{})
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformOutput
type jsiiProxy_TerraformOutput struct {
	jsiiProxy_TerraformElement
}

func (j *jsiiProxy_TerraformOutput) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) DependsOn() *[]ITerraformDependable {
	var returns *[]ITerraformDependable
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) Sensitive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformOutput) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformOutput(scope constructs.Construct, id *string, config *TerraformOutputConfig) TerraformOutput {
	_init_.Initialize()

	j := jsiiProxy_TerraformOutput{}

	_jsii_.Create(
		"cdktf.TerraformOutput",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformOutput_Override(t TerraformOutput, scope constructs.Construct, id *string, config *TerraformOutputConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformOutput",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TerraformOutput) SetDependsOn(val *[]ITerraformDependable) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TerraformOutput) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TerraformOutput) SetSensitive(val *bool) {
	_jsii_.Set(
		j,
		"sensitive",
		val,
	)
}

func (j *jsiiProxy_TerraformOutput) SetValue(val interface{}) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformOutput) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformOutput) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformOutput) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformOutput) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformOutput) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformOutput) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformOutput) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformOutput) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformOutputConfig struct {
	// Experimental.
	DependsOn *[]ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Description *string `json:"description"`
	// Experimental.
	Sensitive *bool `json:"sensitive"`
	// Experimental.
	Value interface{} `json:"value"`
}

// Experimental.
type TerraformProvider interface {
	TerraformElement
	Alias() *string
	SetAlias(val *string)
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	MetaAttributes() *map[string]interface{}
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata
	TerraformProviderSource() *string
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformProvider
type jsiiProxy_TerraformProvider struct {
	jsiiProxy_TerraformElement
}

func (j *jsiiProxy_TerraformProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata {
	var returns *TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformProvider_Override(t TerraformProvider, scope constructs.Construct, id *string, config *TerraformProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformProvider",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TerraformProvider) SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformProvider) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformProvider) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformProvider) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (t *jsiiProxy_TerraformProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformProviderConfig struct {
	// Experimental.
	TerraformResourceType *string `json:"terraformResourceType"`
	// Experimental.
	TerraformGeneratorMetadata *TerraformProviderGeneratorMetadata `json:"terraformGeneratorMetadata"`
	// Experimental.
	TerraformProviderSource *string `json:"terraformProviderSource"`
}

// Experimental.
type TerraformProviderGeneratorMetadata struct {
	// Experimental.
	ProviderName *string `json:"providerName"`
	// Experimental.
	ProviderVersionConstraint *string `json:"providerVersionConstraint"`
}

// Experimental.
type TerraformRemoteState interface {
	TerraformElement
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	FriendlyUniqueId() *string
	RawOverrides() interface{}
	AddOverride(path *string, value interface{})
	Get(output *string) interface{}
	GetBoolean(output *string) *bool
	GetList(output *string) *[]*string
	GetNumber(output *string) *float64
	GetString(output *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformRemoteState
type jsiiProxy_TerraformRemoteState struct {
	jsiiProxy_TerraformElement
}

func (j *jsiiProxy_TerraformRemoteState) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformRemoteState) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformRemoteState) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformRemoteState) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformRemoteState) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformRemoteState_Override(t TerraformRemoteState, scope constructs.Construct, id *string, backend *string, config *DataTerraformRemoteStateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformRemoteState",
		[]interface{}{scope, id, backend, config},
		t,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformRemoteState) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformRemoteState) Get(output *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformRemoteState) GetBoolean(output *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"getBoolean",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformRemoteState) GetList(output *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getList",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformRemoteState) GetNumber(output *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumber",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformRemoteState) GetString(output *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformRemoteState) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformRemoteState) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformRemoteState) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformRemoteState) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformRemoteState) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (t *jsiiProxy_TerraformRemoteState) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformResource interface {
	TerraformElement
	ITerraformDependable
	ITerraformResource
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Lifecycle() *TerraformResourceLifecycle
	SetLifecycle(val *TerraformResourceLifecycle)
	Provider() TerraformProvider
	SetProvider(val TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) *bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformResource
type jsiiProxy_TerraformResource struct {
	jsiiProxy_TerraformElement
	jsiiProxy_ITerraformDependable
	jsiiProxy_ITerraformResource
}

func (j *jsiiProxy_TerraformResource) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Lifecycle() *TerraformResourceLifecycle {
	var returns *TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) Provider() TerraformProvider {
	var returns TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) TerraformGeneratorMetadata() *TerraformProviderGeneratorMetadata {
	var returns *TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformResource(scope constructs.Construct, id *string, config *TerraformResourceConfig) TerraformResource {
	_init_.Initialize()

	j := jsiiProxy_TerraformResource{}

	_jsii_.Create(
		"cdktf.TerraformResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformResource_Override(t TerraformResource, scope constructs.Construct, id *string, config *TerraformResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformResource",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TerraformResource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TerraformResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TerraformResource) SetLifecycle(val *TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TerraformResource) SetProvider(val TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformResource) GetBooleanAttribute(terraformAttribute *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformResource) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformResource) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformResource) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformResource) InterpolationForAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformResource) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformResource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformResource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformResource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (t *jsiiProxy_TerraformResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformResourceConfig struct {
	// Experimental.
	Count *float64 `json:"count"`
	// Experimental.
	DependsOn *[]ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider TerraformProvider `json:"provider"`
	// Experimental.
	TerraformResourceType *string `json:"terraformResourceType"`
	// Experimental.
	TerraformGeneratorMetadata *TerraformProviderGeneratorMetadata `json:"terraformGeneratorMetadata"`
}

// Experimental.
type TerraformResourceLifecycle struct {
	// Experimental.
	CreateBeforeDestroy *bool `json:"createBeforeDestroy"`
	// Experimental.
	IgnoreChanges *[]*string `json:"ignoreChanges"`
	// Experimental.
	PreventDestroy *bool `json:"preventDestroy"`
}

// Experimental.
type TerraformStack interface {
	constructs.Construct
	AddOverride(path *string, value interface{})
	AllocateLogicalId(tfElement interface{}) *string
	AllProviders() *[]TerraformProvider
	GetLogicalId(tfElement interface{}) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformStack
type jsiiProxy_TerraformStack struct {
	internal.Type__constructsConstruct
}

// Experimental.
func NewTerraformStack(scope constructs.Construct, id *string) TerraformStack {
	_init_.Initialize()

	j := jsiiProxy_TerraformStack{}

	_jsii_.Create(
		"cdktf.TerraformStack",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformStack_Override(t TerraformStack, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformStack",
		[]interface{}{scope, id},
		t,
	)
}

// Experimental.
func TerraformStack_IsStack(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.TerraformStack",
		"isStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TerraformStack_Of(construct constructs.IConstruct) TerraformStack {
	_init_.Initialize()

	var returns TerraformStack

	_jsii_.StaticInvoke(
		"cdktf.TerraformStack",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformStack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Returns the naming scheme used to allocate logical IDs.
//
// By default, uses
// the `HashedAddressingScheme` but this method can be overridden to customize
// this behavior.
// Experimental.
func (t *jsiiProxy_TerraformStack) AllocateLogicalId(tfElement interface{}) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"allocateLogicalId",
		[]interface{}{tfElement},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformStack) AllProviders() *[]TerraformProvider {
	var returns *[]TerraformProvider

	_jsii_.Invoke(
		t,
		"allProviders",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformStack) GetLogicalId(tfElement interface{}) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getLogicalId",
		[]interface{}{tfElement},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformStack) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformStack) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformStack) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformStack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformStackMetadata struct {
	// Experimental.
	StackName *string `json:"stackName"`
	// Experimental.
	Version *string `json:"version"`
}

// Experimental.
type TerraformVariable interface {
	TerraformElement
	BooleanValue() *bool
	CdktfStack() TerraformStack
	ConstructNode() constructs.Node
	ConstructNodeMetadata() *map[string]interface{}
	Default() interface{}
	Description() *string
	FriendlyUniqueId() *string
	ListValue() *[]*string
	NumberValue() *float64
	RawOverrides() interface{}
	Sensitive() *bool
	StringValue() *string
	Type() *string
	Value() interface{}
	AddOverride(path *string, value interface{})
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	SynthesizeAttributes() *map[string]interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for TerraformVariable
type jsiiProxy_TerraformVariable struct {
	jsiiProxy_TerraformElement
}

func (j *jsiiProxy_TerraformVariable) BooleanValue() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"booleanValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) CdktfStack() TerraformStack {
	var returns TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) ConstructNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"constructNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Default() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) ListValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) NumberValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Sensitive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TerraformVariable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewTerraformVariable(scope constructs.Construct, id *string, config *TerraformVariableConfig) TerraformVariable {
	_init_.Initialize()

	j := jsiiProxy_TerraformVariable{}

	_jsii_.Create(
		"cdktf.TerraformVariable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Experimental.
func NewTerraformVariable_Override(t TerraformVariable, scope constructs.Construct, id *string, config *TerraformVariableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TerraformVariable",
		[]interface{}{scope, id, config},
		t,
	)
}

// Experimental.
func (t *jsiiProxy_TerraformVariable) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TerraformVariable) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TerraformVariable) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (t *jsiiProxy_TerraformVariable) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (t *jsiiProxy_TerraformVariable) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Experimental.
func (t *jsiiProxy_TerraformVariable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_TerraformVariable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_TerraformVariable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type TerraformVariableConfig struct {
	// Experimental.
	Default interface{} `json:"default"`
	// Experimental.
	Description *string `json:"description"`
	// Experimental.
	Sensitive *bool `json:"sensitive"`
	// The type argument in a variable block allows you to restrict the type of value that will be accepted as the value for a variable.
	//
	// If no type constraint is set then a value of any type is accepted.
	//
	// While type constraints are optional, we recommend specifying them; they serve as easy reminders for users of the module, and allow Terraform to return a helpful error message if the wrong type is used.
	//
	// Type constraints are created from a mixture of type keywords and type constructors. The supported type keywords are:
	//
	// - string
	// - number
	// - bool
	//
	// The type constructors allow you to specify complex types such as collections:
	//
	// - list(\<TYPE\>)
	// - set(\<TYPE\>)
	// - map(\<TYPE\>)
	// - object({\<ATTR NAME\> = \<TYPE\>, ... })
	// - tuple([\<TYPE\>, ...])
	//
	// The keyword any may be used to indicate that any type is acceptable. For more information on the meaning and behavior of these different types, as well as detailed information about automatic conversion of complex types, see {@link https://www.terraform.io/docs/configuration/types.html|Type Constraints}.
	//
	// If both the type and default arguments are specified, the given default value must be convertible to the specified type.
	// Experimental.
	Type *string `json:"type"`
}

// Testing utilities for cdktf applications.
// Experimental.
type Testing interface {
}

// The jsii proxy struct for Testing
type jsiiProxy_Testing struct {
	_ byte // padding
}

// Returns an app for testing with the following properties: - Output directory is a temp dir.
// Experimental.
func Testing_App() App {
	_init_.Initialize()

	var returns App

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"app",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func Testing_EnableFutureFlags(app App) App {
	_init_.Initialize()

	var returns App

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"enableFutureFlags",
		[]interface{}{app},
		&returns,
	)

	return returns
}

// Experimental.
func Testing_StubVersion(app App) App {
	_init_.Initialize()

	var returns App

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"stubVersion",
		[]interface{}{app},
		&returns,
	)

	return returns
}

// Returns the Terraform synthesized JSON.
// Experimental.
func Testing_Synth(stack TerraformStack) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Testing",
		"synth",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

// Represents a special or lazily-evaluated value.
//
// Can be used to delay evaluation of a certain value in case, for example,
// that it requires some context or late-bound data. Can also be used to
// mark values that need special processing at document rendering time.
//
// Tokens can be embedded into strings while retaining their original
// semantics.
// Experimental.
type Token interface {
}

// The jsii proxy struct for Token
type jsiiProxy_Token struct {
	_ byte // padding
}

// Experimental.
func NewToken() Token {
	_init_.Initialize()

	j := jsiiProxy_Token{}

	_jsii_.Create(
		"cdktf.Token",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewToken_Override(t Token) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Token",
		nil, // no parameters
		t,
	)
}

// Return a resolvable representation of the given value.
// Experimental.
func Token_AsAny(value interface{}) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asAny",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Return a reversible list representation of this token.
// Experimental.
func Token_AsList(value interface{}, options *EncodingOptions) *[]*string {
	_init_.Initialize()

	var returns *[]*string

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asList",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Return a reversible number representation of this token.
// Experimental.
func Token_AsNumber(value interface{}) *float64 {
	_init_.Initialize()

	var returns *float64

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Return a reversible string representation of this token.
//
// If the Token is initialized with a literal, the stringified value of the
// literal is returned. Otherwise, a special quoted string representation
// of the Token is returned that can be embedded into other strings.
//
// Strings with quoted Tokens in them can be restored back into
// complex values with the Tokens restored by calling `resolve()`
// on the string.
// Experimental.
func Token_AsString(value interface{}, options *EncodingOptions) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"asString",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

// Returns true if obj represents an unresolved value.
//
// One of these must be true:
//
// - `obj` is an IResolvable
// - `obj` is a string containing at least one encoded `IResolvable`
// - `obj` is either an encoded number or list
//
// This does NOT recurse into lists or objects to see if they
// containing resolvables.
// Experimental.
func Token_IsUnresolved(obj interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.Token",
		"isUnresolved",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Less oft-needed functions to manipulate Tokens.
// Experimental.
type Tokenization interface {
}

// The jsii proxy struct for Tokenization
type jsiiProxy_Tokenization struct {
	_ byte // padding
}

// Experimental.
func NewTokenization() Tokenization {
	_init_.Initialize()

	j := jsiiProxy_Tokenization{}

	_jsii_.Create(
		"cdktf.Tokenization",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTokenization_Override(t Tokenization) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.Tokenization",
		nil, // no parameters
		t,
	)
}

// Return whether the given object is an IResolvable object.
//
// This is different from Token.isUnresolved() which will also check for
// encoded Tokens, whereas this method will only do a type check on the given
// object.
// Experimental.
func Tokenization_IsResolvable(obj interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"isResolvable",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Resolves an object by evaluating all tokens and removing any undefined or empty objects or arrays.
//
// Values can only be primitives, arrays or tokens. Other objects (i.e. with methods) will be rejected.
// Experimental.
func Tokenization_Resolve(obj interface{}, options *ResolveOptions) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"resolve",
		[]interface{}{obj, options},
		&returns,
	)

	return returns
}

// Un-encode a Tokenized value from a list.
// Experimental.
func Tokenization_ReverseList(l *[]*string) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"reverseList",
		[]interface{}{l},
		&returns,
	)

	return returns
}

// Un-encode a Tokenized value from a number.
// Experimental.
func Tokenization_ReverseNumber(n *float64) IResolvable {
	_init_.Initialize()

	var returns IResolvable

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"reverseNumber",
		[]interface{}{n},
		&returns,
	)

	return returns
}

// Un-encode a string potentially containing encoded tokens.
// Experimental.
func Tokenization_ReverseString(s *string) TokenizedStringFragments {
	_init_.Initialize()

	var returns TokenizedStringFragments

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"reverseString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// Stringify a number directly or lazily if it's a Token.
//
// If it is an object (i.e., { Ref: 'SomeLogicalId' }), return it as-is.
// Experimental.
func Tokenization_StringifyNumber(x *float64) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.Tokenization",
		"stringifyNumber",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Fragments of a concatenated string containing stringified Tokens.
// Experimental.
type TokenizedStringFragments interface {
	FirstToken() IResolvable
	FirstValue() interface{}
	Length() *float64
	Tokens() *[]IResolvable
	AddIntrinsic(value interface{})
	AddLiteral(lit interface{})
	AddToken(token IResolvable)
	Join(concat IFragmentConcatenator) interface{}
	MapTokens(mapper ITokenMapper) TokenizedStringFragments
}

// The jsii proxy struct for TokenizedStringFragments
type jsiiProxy_TokenizedStringFragments struct {
	_ byte // padding
}

func (j *jsiiProxy_TokenizedStringFragments) FirstToken() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"firstToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) FirstValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Length() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"length",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenizedStringFragments) Tokens() *[]IResolvable {
	var returns *[]IResolvable
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}


// Experimental.
func NewTokenizedStringFragments() TokenizedStringFragments {
	_init_.Initialize()

	j := jsiiProxy_TokenizedStringFragments{}

	_jsii_.Create(
		"cdktf.TokenizedStringFragments",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTokenizedStringFragments_Override(t TokenizedStringFragments) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.TokenizedStringFragments",
		nil, // no parameters
		t,
	)
}

// Adds an intrinsic fragment.
// Experimental.
func (t *jsiiProxy_TokenizedStringFragments) AddIntrinsic(value interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addIntrinsic",
		[]interface{}{value},
	)
}

// Adds a literal fragment.
// Experimental.
func (t *jsiiProxy_TokenizedStringFragments) AddLiteral(lit interface{}) {
	_jsii_.InvokeVoid(
		t,
		"addLiteral",
		[]interface{}{lit},
	)
}

// Adds a token fragment.
// Experimental.
func (t *jsiiProxy_TokenizedStringFragments) AddToken(token IResolvable) {
	_jsii_.InvokeVoid(
		t,
		"addToken",
		[]interface{}{token},
	)
}

// Combine the string fragments using the given joiner.
//
// If there are any
// Experimental.
func (t *jsiiProxy_TokenizedStringFragments) Join(concat IFragmentConcatenator) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"join",
		[]interface{}{concat},
		&returns,
	)

	return returns
}

// Apply a transformation function to all tokens in the string.
// Experimental.
func (t *jsiiProxy_TokenizedStringFragments) MapTokens(mapper ITokenMapper) TokenizedStringFragments {
	var returns TokenizedStringFragments

	_jsii_.Invoke(
		t,
		"mapTokens",
		[]interface{}{mapper},
		&returns,
	)

	return returns
}

// Experimental.
type VariableType interface {
}

// The jsii proxy struct for VariableType
type jsiiProxy_VariableType struct {
	_ byte // padding
}

// Experimental.
func NewVariableType_Override(v VariableType) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktf.VariableType",
		nil, // no parameters
		v,
	)
}

// Experimental.
func VariableType_List(type_ *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.VariableType",
		"list",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Experimental.
func VariableType_Map(type_ *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.VariableType",
		"map",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Experimental.
func VariableType_Object(attributes *map[string]*string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.VariableType",
		"object",
		[]interface{}{attributes},
		&returns,
	)

	return returns
}

// Experimental.
func VariableType_Set(type_ *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.VariableType",
		"set",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Experimental.
func VariableType_Tuple(elements ...*string) *string {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range elements {
		args = append(args, a)
	}

	var returns *string

	_jsii_.StaticInvoke(
		"cdktf.VariableType",
		"tuple",
		args,
		&returns,
	)

	return returns
}

func VariableType_ANY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"ANY",
		&returns,
	)
	return returns
}

func VariableType_BOOL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"BOOL",
		&returns,
	)
	return returns
}

func VariableType_LIST() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"LIST",
		&returns,
	)
	return returns
}

func VariableType_LIST_BOOL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"LIST_BOOL",
		&returns,
	)
	return returns
}

func VariableType_LIST_NUMBER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"LIST_NUMBER",
		&returns,
	)
	return returns
}

func VariableType_LIST_STRING() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"LIST_STRING",
		&returns,
	)
	return returns
}

func VariableType_MAP() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"MAP",
		&returns,
	)
	return returns
}

func VariableType_MAP_BOOL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"MAP_BOOL",
		&returns,
	)
	return returns
}

func VariableType_MAP_NUMBER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"MAP_NUMBER",
		&returns,
	)
	return returns
}

func VariableType_MAP_STRING() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"MAP_STRING",
		&returns,
	)
	return returns
}

func VariableType_NUMBER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"NUMBER",
		&returns,
	)
	return returns
}

func VariableType_SET() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"SET",
		&returns,
	)
	return returns
}

func VariableType_SET_BOOL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"SET_BOOL",
		&returns,
	)
	return returns
}

func VariableType_SET_NUMBER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"SET_NUMBER",
		&returns,
	)
	return returns
}

func VariableType_SET_STRING() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"SET_STRING",
		&returns,
	)
	return returns
}

func VariableType_STRING() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktf.VariableType",
		"STRING",
		&returns,
	)
	return returns
}

