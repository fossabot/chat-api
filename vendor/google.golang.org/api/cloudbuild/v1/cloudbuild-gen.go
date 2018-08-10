// Package cloudbuild provides access to the Cloud Build API.
//
// See https://cloud.google.com/cloud-build/docs/
//
// Usage example:
//
//   import "google.golang.org/api/cloudbuild/v1"
//   ...
//   cloudbuildService, err := cloudbuild.New(oauthHttpClient)
package cloudbuild // import "google.golang.org/api/cloudbuild/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "cloudbuild:v1"
const apiName = "cloudbuild"
const apiVersion = "v1"
const basePath = "https://cloudbuild.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Operations = NewOperationsService(s)
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Operations *OperationsService

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Builds = NewProjectsBuildsService(s)
	rs.Triggers = NewProjectsTriggersService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Builds *ProjectsBuildsService

	Triggers *ProjectsTriggersService
}

func NewProjectsBuildsService(s *Service) *ProjectsBuildsService {
	rs := &ProjectsBuildsService{s: s}
	return rs
}

type ProjectsBuildsService struct {
	s *Service
}

func NewProjectsTriggersService(s *Service) *ProjectsTriggersService {
	rs := &ProjectsTriggersService{s: s}
	return rs
}

type ProjectsTriggersService struct {
	s *Service
}

// ArtifactObjects: Files in the workspace to upload to Cloud Storage
// upon successful
// completion of all build steps.
type ArtifactObjects struct {
	// Location: Cloud Storage bucket and optional object path, in the
	// form
	// "gs://bucket/path/to/somewhere/". (see [Bucket
	// Name
	// Requirements](https://cloud.google.com/storage/docs/bucket-naming
	// #requirements)).
	//
	// Files in the workspace matching any path pattern will be uploaded
	// to
	// Cloud Storage with this location as a prefix.
	Location string `json:"location,omitempty"`

	// Paths: Path globs used to match files in the build's workspace.
	Paths []string `json:"paths,omitempty"`

	// Timing: Output only. Stores timing information for pushing all
	// artifact objects.
	Timing *TimeSpan `json:"timing,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Location") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Location") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ArtifactObjects) MarshalJSON() ([]byte, error) {
	type NoMethod ArtifactObjects
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ArtifactResult: An artifact that was uploaded during a build. This
// is a single record in the artifact manifest JSON file.
type ArtifactResult struct {
	// FileHash: The file hash of the artifact.
	FileHash []*FileHashes `json:"fileHash,omitempty"`

	// Location: The path of an artifact in a Google Cloud Storage bucket,
	// with the
	// generation number. For
	// example,
	// `gs://mybucket/path/to/output.jar#generation`.
	Location string `json:"location,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FileHash") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FileHash") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ArtifactResult) MarshalJSON() ([]byte, error) {
	type NoMethod ArtifactResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Artifacts: Artifacts produced by a build that should be uploaded
// upon
// successful completion of all build steps.
type Artifacts struct {
	// Images: A list of images to be pushed upon the successful completion
	// of all build
	// steps.
	//
	// The images will be pushed using the builder service account's
	// credentials.
	//
	// The digests of the pushed images will be stored in the Build
	// resource's
	// results field.
	//
	// If any of the images fail to be pushed, the build is marked FAILURE.
	Images []string `json:"images,omitempty"`

	// Objects: A list of objects to be uploaded to Cloud Storage upon
	// successful
	// completion of all build steps.
	//
	// Files in the workspace matching specified paths globs will be
	// uploaded to
	// the specified Cloud Storage location using the builder service
	// account's
	// credentials.
	//
	// The location and generation of the uploaded objects will be stored in
	// the
	// Build resource's results field.
	//
	// If any objects fail to be pushed, the build is marked FAILURE.
	Objects *ArtifactObjects `json:"objects,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Images") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Images") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Artifacts) MarshalJSON() ([]byte, error) {
	type NoMethod Artifacts
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Build: A build resource in the Cloud Build API.
//
// At a high level, a `Build` describes where to find source code, how
// to build
// it (for example, the builder image to run on the source), and where
// to store
// the built artifacts.
//
// Fields can include the following variables, which will be expanded
// when the
// build is created:
//
// - $PROJECT_ID: the project ID of the build.
// - $BUILD_ID: the autogenerated ID of the build.
// - $REPO_NAME: the source repository name specified by RepoSource.
// - $BRANCH_NAME: the branch name specified by RepoSource.
// - $TAG_NAME: the tag name specified by RepoSource.
// - $REVISION_ID or $COMMIT_SHA: the commit SHA specified by RepoSource
// or
//   resolved from the specified branch or tag.
// - $SHORT_SHA: first 7 characters of $REVISION_ID or $COMMIT_SHA.
type Build struct {
	// Artifacts: Artifacts produced by the build that should be uploaded
	// upon
	// successful completion of all build steps.
	Artifacts *Artifacts `json:"artifacts,omitempty"`

	// BuildTriggerId: Output only. The ID of the `BuildTrigger` that
	// triggered this build, if it
	// was triggered automatically.
	BuildTriggerId string `json:"buildTriggerId,omitempty"`

	// CreateTime: Output only. Time at which the request to create the
	// build was received.
	CreateTime string `json:"createTime,omitempty"`

	// FinishTime: Output only. Time at which execution of the build was
	// finished.
	//
	// The difference between finish_time and start_time is the duration of
	// the
	// build's execution.
	FinishTime string `json:"finishTime,omitempty"`

	// Id: Output only. Unique identifier of the build.
	Id string `json:"id,omitempty"`

	// Images: A list of images to be pushed upon the successful completion
	// of all build
	// steps.
	//
	// The images are pushed using the builder service account's
	// credentials.
	//
	// The digests of the pushed images will be stored in the `Build`
	// resource's
	// results field.
	//
	// If any of the images fail to be pushed, the build status is
	// marked
	// `FAILURE`.
	Images []string `json:"images,omitempty"`

	// LogUrl: Output only. URL to logs for this build in Google Cloud
	// Console.
	LogUrl string `json:"logUrl,omitempty"`

	// LogsBucket: Google Cloud Storage bucket where logs should be written
	// (see
	// [Bucket
	// Name
	// Requirements](https://cloud.google.com/storage/docs/bucket-naming
	// #requirements)).
	// Logs file names will be of the format
	// `${logs_bucket}/log-${build_id}.txt`.
	LogsBucket string `json:"logsBucket,omitempty"`

	// Options: Special options for this build.
	Options *BuildOptions `json:"options,omitempty"`

	// ProjectId: Output only. ID of the project.
	ProjectId string `json:"projectId,omitempty"`

	// Results: Output only. Results of the build.
	Results *Results `json:"results,omitempty"`

	// Secrets: Secrets to decrypt using Cloud Key Management Service.
	Secrets []*Secret `json:"secrets,omitempty"`

	// Source: The location of the source files to build.
	Source *Source `json:"source,omitempty"`

	// SourceProvenance: Output only. A permanent fixed identifier for
	// source.
	SourceProvenance *SourceProvenance `json:"sourceProvenance,omitempty"`

	// StartTime: Output only. Time at which execution of the build was
	// started.
	StartTime string `json:"startTime,omitempty"`

	// Status: Output only. Status of the build.
	//
	// Possible values:
	//   "STATUS_UNKNOWN" - Status of the build is unknown.
	//   "QUEUED" - Build or step is queued; work has not yet begun.
	//   "WORKING" - Build or step is being executed.
	//   "SUCCESS" - Build or step finished successfully.
	//   "FAILURE" - Build or step failed to complete successfully.
	//   "INTERNAL_ERROR" - Build or step failed due to an internal cause.
	//   "TIMEOUT" - Build or step took longer than was allowed.
	//   "CANCELLED" - Build or step was canceled by a user.
	Status string `json:"status,omitempty"`

	// StatusDetail: Output only. Customer-readable message about the
	// current status.
	StatusDetail string `json:"statusDetail,omitempty"`

	// Steps: Required. The operations to be performed on the workspace.
	Steps []*BuildStep `json:"steps,omitempty"`

	// Substitutions: Substitutions data for `Build` resource.
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// Tags: Tags for annotation of a `Build`. These are not docker tags.
	Tags []string `json:"tags,omitempty"`

	// Timeout: Amount of time that this build should be allowed to run, to
	// second
	// granularity. If this amount of time elapses, work on the build will
	// cease
	// and the build status will be `TIMEOUT`.
	//
	// Default time is ten minutes.
	Timeout string `json:"timeout,omitempty"`

	// Timing: Output only. Stores timing information for phases of the
	// build. Valid keys
	// are:
	//
	// * BUILD: time to execute all build steps
	// * PUSH: time to push all specified images.
	// * FETCHSOURCE: time to fetch source.
	//
	// If the build does not specify source or images,
	// these keys will not be included.
	Timing map[string]TimeSpan `json:"timing,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Artifacts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Artifacts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Build) MarshalJSON() ([]byte, error) {
	type NoMethod Build
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BuildOperationMetadata: Metadata for build operations.
type BuildOperationMetadata struct {
	// Build: The build that the operation is tracking.
	Build *Build `json:"build,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Build") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Build") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BuildOperationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod BuildOperationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BuildOptions: Optional arguments to enable specific features of
// builds.
type BuildOptions struct {
	// DiskSizeGb: Requested disk size for the VM that runs the build. Note
	// that this is *NOT*
	// "disk free"; some of the space will be used by the operating system
	// and
	// build utilities. Also note that this is the minimum disk size that
	// will be
	// allocated for the build -- the build may run with a larger disk
	// than
	// requested. At present, the maximum disk size is 1000GB; builds that
	// request
	// more than the maximum are rejected with an error.
	DiskSizeGb int64 `json:"diskSizeGb,omitempty,string"`

	// LogStreamingOption: Option to define build log streaming behavior to
	// Google Cloud
	// Storage.
	//
	// Possible values:
	//   "STREAM_DEFAULT" - Service may automatically determine build log
	// streaming behavior.
	//   "STREAM_ON" - Build logs should be streamed to Google Cloud
	// Storage.
	//   "STREAM_OFF" - Build logs should not be streamed to Google Cloud
	// Storage; they will be
	// written when the build is completed.
	LogStreamingOption string `json:"logStreamingOption,omitempty"`

	// MachineType: Compute Engine machine type on which to run the build.
	//
	// Possible values:
	//   "UNSPECIFIED" - Standard machine type.
	//   "N1_HIGHCPU_8" - Highcpu machine with 8 CPUs.
	//   "N1_HIGHCPU_32" - Highcpu machine with 32 CPUs.
	MachineType string `json:"machineType,omitempty"`

	// RequestedVerifyOption: Requested verifiability options.
	//
	// Possible values:
	//   "NOT_VERIFIED" - Not a verifiable build. (default)
	//   "VERIFIED" - Verified build.
	RequestedVerifyOption string `json:"requestedVerifyOption,omitempty"`

	// SourceProvenanceHash: Requested hash for SourceProvenance.
	//
	// Possible values:
	//   "NONE" - No hash requested.
	//   "SHA256" - Use a sha256 hash.
	//   "MD5" - Use a md5 hash.
	SourceProvenanceHash []string `json:"sourceProvenanceHash,omitempty"`

	// SubstitutionOption: Option to specify behavior when there is an error
	// in the substitution
	// checks.
	//
	// Possible values:
	//   "MUST_MATCH" - Fails the build if error in substitutions checks,
	// like missing
	// a substitution in the template or in the map.
	//   "ALLOW_LOOSE" - Do not fail the build if error in substitutions
	// checks.
	SubstitutionOption string `json:"substitutionOption,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DiskSizeGb") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DiskSizeGb") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BuildOptions) MarshalJSON() ([]byte, error) {
	type NoMethod BuildOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BuildStep: A step in the build pipeline.
type BuildStep struct {
	// Args: A list of arguments that will be presented to the step when it
	// is started.
	//
	// If the image used to run the step's container has an entrypoint, the
	// `args`
	// are used as arguments to that entrypoint. If the image does not
	// define
	// an entrypoint, the first element in args is used as the
	// entrypoint,
	// and the remainder will be used as arguments.
	Args []string `json:"args,omitempty"`

	// Dir: Working directory to use when running this step's container.
	//
	// If this value is a relative path, it is relative to the build's
	// working
	// directory. If this value is absolute, it may be outside the build's
	// working
	// directory, in which case the contents of the path may not be
	// persisted
	// across build step executions, unless a `volume` for that path is
	// specified.
	//
	// If the build specifies a `RepoSource` with `dir` and a step with a
	// `dir`,
	// which specifies an absolute path, the `RepoSource` `dir` is ignored
	// for
	// the step's execution.
	Dir string `json:"dir,omitempty"`

	// Entrypoint: Entrypoint to be used instead of the build step image's
	// default entrypoint.
	// If unset, the image's default entrypoint is used.
	Entrypoint string `json:"entrypoint,omitempty"`

	// Env: A list of environment variable definitions to be used when
	// running a step.
	//
	// The elements are of the form "KEY=VALUE" for the environment variable
	// "KEY"
	// being given the value "VALUE".
	Env []string `json:"env,omitempty"`

	// Id: Unique identifier for this build step, used in `wait_for`
	// to
	// reference this build step as a dependency.
	Id string `json:"id,omitempty"`

	// Name: Required. The name of the container image that will run this
	// particular
	// build step.
	//
	// If the image is available in the host's Docker daemon's cache,
	// it
	// will be run directly. If not, the host will attempt to pull the
	// image
	// first, using the builder service account's credentials if
	// necessary.
	//
	// The Docker daemon's cache will already have the latest versions of
	// all of
	// the officially supported build
	// steps
	// ([https://github.com/GoogleCloudPlatform/cloud-builders](https:/
	// /github.com/GoogleCloudPlatform/cloud-builders)).
	// The Docker daemon will also have cached many of the layers for some
	// popular
	// images, like "ubuntu", "debian", but they will be refreshed at the
	// time you
	// attempt to use them.
	//
	// If you built an image in a previous build step, it will be stored in
	// the
	// host's Docker daemon's cache and is available to use as the name for
	// a
	// later build step.
	Name string `json:"name,omitempty"`

	// SecretEnv: A list of environment variables which are encrypted using
	// a Cloud Key
	// Management Service crypto key. These values must be specified in
	// the
	// build's `Secret`.
	SecretEnv []string `json:"secretEnv,omitempty"`

	// Status: Output only. Status of the build step. At this time, build
	// step status is
	// only updated on build completion; step status is not updated in
	// real-time
	// as the build progresses.
	//
	// Possible values:
	//   "STATUS_UNKNOWN" - Status of the build is unknown.
	//   "QUEUED" - Build or step is queued; work has not yet begun.
	//   "WORKING" - Build or step is being executed.
	//   "SUCCESS" - Build or step finished successfully.
	//   "FAILURE" - Build or step failed to complete successfully.
	//   "INTERNAL_ERROR" - Build or step failed due to an internal cause.
	//   "TIMEOUT" - Build or step took longer than was allowed.
	//   "CANCELLED" - Build or step was canceled by a user.
	Status string `json:"status,omitempty"`

	// Timeout: Time limit for executing this build step. If not defined,
	// the step has no
	// time limit and will be allowed to continue to run until either it
	// completes
	// or the build itself times out.
	Timeout string `json:"timeout,omitempty"`

	// Timing: Output only. Stores timing information for executing this
	// build step.
	Timing *TimeSpan `json:"timing,omitempty"`

	// Volumes: List of volumes to mount into the build step.
	//
	// Each volume will be created as an empty volume prior to execution of
	// the
	// build step. Upon completion of the build, volumes and their contents
	// will
	// be discarded.
	//
	// Using a named volume in only one step is not valid as it is
	// indicative
	// of a mis-configured build request.
	Volumes []*Volume `json:"volumes,omitempty"`

	// WaitFor: The ID(s) of the step(s) that this build step depends
	// on.
	// This build step will not start until all the build steps in
	// `wait_for`
	// have completed successfully. If `wait_for` is empty, this build step
	// will
	// start when all previous build steps in the `Build.Steps` list
	// have
	// completed successfully.
	WaitFor []string `json:"waitFor,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Args") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Args") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BuildStep) MarshalJSON() ([]byte, error) {
	type NoMethod BuildStep
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BuildTrigger: Configuration for an automated build in response to
// source repository
// changes.
type BuildTrigger struct {
	// Build: Contents of the build template.
	Build *Build `json:"build,omitempty"`

	// CreateTime: Output only. Time when the trigger was created.
	CreateTime string `json:"createTime,omitempty"`

	// Description: Human-readable description of this trigger.
	Description string `json:"description,omitempty"`

	// Disabled: If true, the trigger will never result in a build.
	Disabled bool `json:"disabled,omitempty"`

	// Filename: Path, from the source root, to a file whose contents is
	// used for the
	// template.
	Filename string `json:"filename,omitempty"`

	// Id: Output only. Unique identifier of the trigger.
	Id string `json:"id,omitempty"`

	// Substitutions: Substitutions data for Build resource.
	Substitutions map[string]string `json:"substitutions,omitempty"`

	// TriggerTemplate: Template describing the types of source changes to
	// trigger a build.
	//
	// Branch and tag names in trigger templates are interpreted as
	// regular
	// expressions. Any branch or tag change that matches that regular
	// expression
	// will trigger a build.
	TriggerTemplate *RepoSource `json:"triggerTemplate,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Build") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Build") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BuildTrigger) MarshalJSON() ([]byte, error) {
	type NoMethod BuildTrigger
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BuiltImage: An image built by the pipeline.
type BuiltImage struct {
	// Digest: Docker Registry 2.0 digest.
	Digest string `json:"digest,omitempty"`

	// Name: Name used to push the container image to Google Container
	// Registry, as
	// presented to `docker push`.
	Name string `json:"name,omitempty"`

	// PushTiming: Output only. Stores timing information for pushing the
	// specified image.
	PushTiming *TimeSpan `json:"pushTiming,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Digest") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Digest") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BuiltImage) MarshalJSON() ([]byte, error) {
	type NoMethod BuiltImage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CancelBuildRequest: Request to cancel an ongoing build.
type CancelBuildRequest struct {
}

// CancelOperationRequest: The request message for
// Operations.CancelOperation.
type CancelOperationRequest struct {
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// FileHashes: Container message for hashes of byte content of files,
// used in
// SourceProvenance messages to verify integrity of source input to the
// build.
type FileHashes struct {
	// FileHash: Collection of file hashes.
	FileHash []*Hash `json:"fileHash,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FileHash") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FileHash") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FileHashes) MarshalJSON() ([]byte, error) {
	type NoMethod FileHashes
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Hash: Container message for hash values.
type Hash struct {
	// Type: The type of hash that was performed.
	//
	// Possible values:
	//   "NONE" - No hash requested.
	//   "SHA256" - Use a sha256 hash.
	//   "MD5" - Use a md5 hash.
	Type string `json:"type,omitempty"`

	// Value: The hash value.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Hash) MarshalJSON() ([]byte, error) {
	type NoMethod Hash
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListBuildTriggersResponse: Response containing existing
// `BuildTriggers`.
type ListBuildTriggersResponse struct {
	// Triggers: `BuildTriggers` for the project, sorted by `create_time`
	// descending.
	Triggers []*BuildTrigger `json:"triggers,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Triggers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Triggers") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListBuildTriggersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListBuildTriggersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListBuildsResponse: Response including listed builds.
type ListBuildsResponse struct {
	// Builds: Builds will be sorted by `create_time`, descending.
	Builds []*Build `json:"builds,omitempty"`

	// NextPageToken: Token to receive the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Builds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Builds") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListBuildsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListBuildsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListOperationsResponse: The response message for
// Operations.ListOperations.
type ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*Operation `json:"operations,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListOperationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListOperationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a
// network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RepoSource: Location of the source in a Google Cloud Source
// Repository.
type RepoSource struct {
	// BranchName: Name of the branch to build.
	BranchName string `json:"branchName,omitempty"`

	// CommitSha: Explicit commit SHA to build.
	CommitSha string `json:"commitSha,omitempty"`

	// Dir: Directory, relative to the source root, in which to run the
	// build.
	//
	// This must be a relative path. If a step's `dir` is specified and is
	// an
	// absolute path, this value is ignored for that step's execution.
	Dir string `json:"dir,omitempty"`

	// ProjectId: ID of the project that owns the Cloud Source Repository.
	// If omitted, the
	// project ID requesting the build is assumed.
	ProjectId string `json:"projectId,omitempty"`

	// RepoName: Name of the Cloud Source Repository. If omitted, the name
	// "default" is
	// assumed.
	RepoName string `json:"repoName,omitempty"`

	// TagName: Name of the tag to build.
	TagName string `json:"tagName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BranchName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BranchName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RepoSource) MarshalJSON() ([]byte, error) {
	type NoMethod RepoSource
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Results: Artifacts created by the build pipeline.
type Results struct {
	// ArtifactManifest: Path to the artifact manifest. Only populated when
	// artifacts are uploaded.
	ArtifactManifest string `json:"artifactManifest,omitempty"`

	// BuildStepImages: List of build step digests, in the order
	// corresponding to build step
	// indices.
	BuildStepImages []string `json:"buildStepImages,omitempty"`

	// BuildStepOutputs: List of build step outputs, produced by builder
	// images, in the order
	// corresponding to build step indices.
	//
	// Builders can produce this output by writing to
	// `$BUILDER_OUTPUT/output`.
	// Only the first 4KB of data is stored.
	BuildStepOutputs []string `json:"buildStepOutputs,omitempty"`

	// Images: Container images that were built as a part of the build.
	Images []*BuiltImage `json:"images,omitempty"`

	// NumArtifacts: Number of artifacts uploaded. Only populated when
	// artifacts are uploaded.
	NumArtifacts int64 `json:"numArtifacts,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "ArtifactManifest") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ArtifactManifest") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Results) MarshalJSON() ([]byte, error) {
	type NoMethod Results
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RetryBuildRequest: Specifies a build to retry.
type RetryBuildRequest struct {
}

// Secret: Pairs a set of secret environment variables containing
// encrypted
// values with the Cloud KMS key to use to decrypt the value.
type Secret struct {
	// KmsKeyName: Cloud KMS key name to use to decrypt these envs.
	KmsKeyName string `json:"kmsKeyName,omitempty"`

	// SecretEnv: Map of environment variable name to its encrypted
	// value.
	//
	// Secret environment variables must be unique across all of a
	// build's
	// secrets, and must be used by at least one build step. Values can be
	// at most
	// 1 KB in size. There can be at most ten secret values across all of
	// a
	// build's secrets.
	SecretEnv map[string]string `json:"secretEnv,omitempty"`

	// ForceSendFields is a list of field names (e.g. "KmsKeyName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "KmsKeyName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Secret) MarshalJSON() ([]byte, error) {
	type NoMethod Secret
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Source: Location of the source in a supported storage service.
type Source struct {
	// RepoSource: If provided, get the source from this location in a Cloud
	// Source
	// Repository.
	RepoSource *RepoSource `json:"repoSource,omitempty"`

	// StorageSource: If provided, get the source from this location in
	// Google Cloud Storage.
	StorageSource *StorageSource `json:"storageSource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "RepoSource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "RepoSource") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Source) MarshalJSON() ([]byte, error) {
	type NoMethod Source
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SourceProvenance: Provenance of the source. Ways to find the original
// source, or verify that
// some source was used for this build.
type SourceProvenance struct {
	// FileHashes: Output only. Hash(es) of the build source, which can be
	// used to verify that
	// the originalsource integrity was maintained in the build. Note
	// that
	// `FileHashes` willonly be populated if `BuildOptions` has requested
	// a
	// `SourceProvenanceHash`.
	//
	// The keys to this map are file paths used as build source and the
	// values
	// contain the hash values for those files.
	//
	// If the build source came in a single package such as a gzipped
	// tarfile
	// (`.tar.gz`), the `FileHash` will be for the single path to that file.
	FileHashes map[string]FileHashes `json:"fileHashes,omitempty"`

	// ResolvedRepoSource: A copy of the build's `source.repo_source`, if
	// exists, with any
	// revisions resolved.
	ResolvedRepoSource *RepoSource `json:"resolvedRepoSource,omitempty"`

	// ResolvedStorageSource: A copy of the build's `source.storage_source`,
	// if exists, with any
	// generations resolved.
	ResolvedStorageSource *StorageSource `json:"resolvedStorageSource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FileHashes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FileHashes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SourceProvenance) MarshalJSON() ([]byte, error) {
	type NoMethod SourceProvenance
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StorageSource: Location of the source in an archive file in Google
// Cloud Storage.
type StorageSource struct {
	// Bucket: Google Cloud Storage bucket containing the source
	// (see
	// [Bucket
	// Name
	// Requirements](https://cloud.google.com/storage/docs/bucket-naming
	// #requirements)).
	Bucket string `json:"bucket,omitempty"`

	// Generation: Google Cloud Storage generation for the object. If the
	// generation is
	// omitted, the latest generation will be used.
	Generation int64 `json:"generation,omitempty,string"`

	// Object: Google Cloud Storage object containing the source.
	//
	// This object must be a gzipped archive file (`.tar.gz`) containing
	// source to
	// build.
	Object string `json:"object,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bucket") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bucket") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StorageSource) MarshalJSON() ([]byte, error) {
	type NoMethod StorageSource
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimeSpan: Start and end times for a build execution phase.
type TimeSpan struct {
	// EndTime: End of time span.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: Start of time span.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimeSpan) MarshalJSON() ([]byte, error) {
	type NoMethod TimeSpan
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Volume: Volume describes a Docker container volume which is mounted
// into build steps
// in order to persist files across build step execution.
type Volume struct {
	// Name: Name of the volume to mount.
	//
	// Volume names must be unique per build step and must be valid names
	// for
	// Docker volumes. Each named volume must be used by at least two build
	// steps.
	Name string `json:"name,omitempty"`

	// Path: Path at which to mount the volume.
	//
	// Paths must be absolute and cannot conflict with other volume paths on
	// the
	// same build step or with certain reserved volume paths.
	Path string `json:"path,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Volume) MarshalJSON() ([]byte, error) {
	type NoMethod Volume
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "cloudbuild.operations.cancel":

type OperationsCancelCall struct {
	s                      *Service
	name                   string
	canceloperationrequest *CancelOperationRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Cancel: Starts asynchronous cancellation on a long-running operation.
//  The server
// makes a best effort to cancel the operation, but success is
// not
// guaranteed.  If the server doesn't support this method, it
// returns
// `google.rpc.Code.UNIMPLEMENTED`.  Clients can
// use
// Operations.GetOperation or
// other methods to check whether the cancellation succeeded or whether
// the
// operation completed despite cancellation. On successful
// cancellation,
// the operation is not deleted; instead, it becomes an operation
// with
// an Operation.error value with a google.rpc.Status.code of
// 1,
// corresponding to `Code.CANCELLED`.
func (r *OperationsService) Cancel(name string, canceloperationrequest *CancelOperationRequest) *OperationsCancelCall {
	c := &OperationsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.canceloperationrequest = canceloperationrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsCancelCall) Fields(s ...googleapi.Field) *OperationsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsCancelCall) Context(ctx context.Context) *OperationsCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OperationsCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OperationsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.canceloperationrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.operations.cancel" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OperationsCancelCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts asynchronous cancellation on a long-running operation.  The server\nmakes a best effort to cancel the operation, but success is not\nguaranteed.  If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.  Clients can use\nOperations.GetOperation or\nother methods to check whether the cancellation succeeded or whether the\noperation completed despite cancellation. On successful cancellation,\nthe operation is not deleted; instead, it becomes an operation with\nan Operation.error value with a google.rpc.Status.code of 1,\ncorresponding to `Code.CANCELLED`.",
	//   "flatPath": "v1/operations/{operationsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "cloudbuild.operations.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be cancelled.",
	//       "location": "path",
	//       "pattern": "^operations/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:cancel",
	//   "request": {
	//     "$ref": "CancelOperationRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.operations.get":

type OperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *OperationsService) Get(name string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OperationsGetCall) IfNoneMatch(entityTag string) *OperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsGetCall) Context(ctx context.Context) *OperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v1/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudbuild.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^operations/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.operations.list":

type OperationsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists operations that match the specified filter in the
// request. If the
// server doesn't support this method, it returns
// `UNIMPLEMENTED`.
//
// NOTE: the `name` binding allows API services to override the
// binding
// to use different resource name schemes, such as `users/*/operations`.
// To
// override the binding, API services can add a binding such
// as
// "/v1/{name=users/*}/operations" to their service configuration.
// For backwards compatibility, the default name includes the
// operations
// collection id, however overriding users must ensure the name
// binding
// is the parent resource, without the operations collection id.
func (r *OperationsService) List(name string) *OperationsListCall {
	c := &OperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *OperationsListCall) Filter(filter string) *OperationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *OperationsListCall) PageSize(pageSize int64) *OperationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *OperationsListCall) PageToken(pageToken string) *OperationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsListCall) Fields(s ...googleapi.Field) *OperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OperationsListCall) IfNoneMatch(entityTag string) *OperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsListCall) Context(ctx context.Context) *OperationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OperationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OperationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.operations.list" call.
// Exactly one of *ListOperationsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListOperationsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OperationsListCall) Do(opts ...googleapi.CallOption) (*ListOperationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListOperationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists operations that match the specified filter in the request. If the\nserver doesn't support this method, it returns `UNIMPLEMENTED`.\n\nNOTE: the `name` binding allows API services to override the binding\nto use different resource name schemes, such as `users/*/operations`. To\noverride the binding, API services can add a binding such as\n`\"/v1/{name=users/*}/operations\"` to their service configuration.\nFor backwards compatibility, the default name includes the operations\ncollection id, however overriding users must ensure the name binding\nis the parent resource, without the operations collection id.",
	//   "flatPath": "v1/operations",
	//   "httpMethod": "GET",
	//   "id": "cloudbuild.operations.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The standard list filter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name of the operation's parent resource.",
	//       "location": "path",
	//       "pattern": "^operations$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The standard list page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard list page token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OperationsListCall) Pages(ctx context.Context, f func(*ListOperationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudbuild.projects.builds.cancel":

type ProjectsBuildsCancelCall struct {
	s                  *Service
	projectId          string
	id                 string
	cancelbuildrequest *CancelBuildRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// Cancel: Cancels a build in progress.
func (r *ProjectsBuildsService) Cancel(projectId string, id string, cancelbuildrequest *CancelBuildRequest) *ProjectsBuildsCancelCall {
	c := &ProjectsBuildsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.id = id
	c.cancelbuildrequest = cancelbuildrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBuildsCancelCall) Fields(s ...googleapi.Field) *ProjectsBuildsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBuildsCancelCall) Context(ctx context.Context) *ProjectsBuildsCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBuildsCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBuildsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cancelbuildrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/builds/{id}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"id":        c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.builds.cancel" call.
// Exactly one of *Build or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Build.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsBuildsCancelCall) Do(opts ...googleapi.CallOption) (*Build, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Build{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancels a build in progress.",
	//   "flatPath": "v1/projects/{projectId}/builds/{id}:cancel",
	//   "httpMethod": "POST",
	//   "id": "cloudbuild.projects.builds.cancel",
	//   "parameterOrder": [
	//     "projectId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "ID of the build.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/builds/{id}:cancel",
	//   "request": {
	//     "$ref": "CancelBuildRequest"
	//   },
	//   "response": {
	//     "$ref": "Build"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.builds.create":

type ProjectsBuildsCreateCall struct {
	s          *Service
	projectId  string
	build      *Build
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Starts a build with the specified configuration.
//
// This method returns a long-running `Operation`, which includes the
// build
// ID. Pass the build ID to `GetBuild` to determine the build status
// (such as
// `SUCCESS` or `FAILURE`).
func (r *ProjectsBuildsService) Create(projectId string, build *Build) *ProjectsBuildsCreateCall {
	c := &ProjectsBuildsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.build = build
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBuildsCreateCall) Fields(s ...googleapi.Field) *ProjectsBuildsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBuildsCreateCall) Context(ctx context.Context) *ProjectsBuildsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBuildsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBuildsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.build)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/builds")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.builds.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsBuildsCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts a build with the specified configuration.\n\nThis method returns a long-running `Operation`, which includes the build\nID. Pass the build ID to `GetBuild` to determine the build status (such as\n`SUCCESS` or `FAILURE`).",
	//   "flatPath": "v1/projects/{projectId}/builds",
	//   "httpMethod": "POST",
	//   "id": "cloudbuild.projects.builds.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/builds",
	//   "request": {
	//     "$ref": "Build"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.builds.get":

type ProjectsBuildsGetCall struct {
	s            *Service
	projectId    string
	id           string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns information about a previously requested build.
//
// The `Build` that is returned includes its status (such as
// `SUCCESS`,
// `FAILURE`, or `WORKING`), and timing information.
func (r *ProjectsBuildsService) Get(projectId string, id string) *ProjectsBuildsGetCall {
	c := &ProjectsBuildsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBuildsGetCall) Fields(s ...googleapi.Field) *ProjectsBuildsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBuildsGetCall) IfNoneMatch(entityTag string) *ProjectsBuildsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBuildsGetCall) Context(ctx context.Context) *ProjectsBuildsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBuildsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBuildsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/builds/{id}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"id":        c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.builds.get" call.
// Exactly one of *Build or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Build.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsBuildsGetCall) Do(opts ...googleapi.CallOption) (*Build, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Build{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns information about a previously requested build.\n\nThe `Build` that is returned includes its status (such as `SUCCESS`,\n`FAILURE`, or `WORKING`), and timing information.",
	//   "flatPath": "v1/projects/{projectId}/builds/{id}",
	//   "httpMethod": "GET",
	//   "id": "cloudbuild.projects.builds.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "ID of the build.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/builds/{id}",
	//   "response": {
	//     "$ref": "Build"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.builds.list":

type ProjectsBuildsListCall struct {
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists previously requested builds.
//
// Previously requested builds may still be in-progress, or may have
// finished
// successfully or unsuccessfully.
func (r *ProjectsBuildsService) List(projectId string) *ProjectsBuildsListCall {
	c := &ProjectsBuildsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// Filter sets the optional parameter "filter": The raw filter text to
// constrain the results.
func (c *ProjectsBuildsListCall) Filter(filter string) *ProjectsBuildsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": Number of results to
// return in the list.
func (c *ProjectsBuildsListCall) PageSize(pageSize int64) *ProjectsBuildsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to provide
// to skip to a particular spot in the list.
func (c *ProjectsBuildsListCall) PageToken(pageToken string) *ProjectsBuildsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBuildsListCall) Fields(s ...googleapi.Field) *ProjectsBuildsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsBuildsListCall) IfNoneMatch(entityTag string) *ProjectsBuildsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBuildsListCall) Context(ctx context.Context) *ProjectsBuildsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBuildsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBuildsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/builds")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.builds.list" call.
// Exactly one of *ListBuildsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListBuildsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsBuildsListCall) Do(opts ...googleapi.CallOption) (*ListBuildsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListBuildsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists previously requested builds.\n\nPreviously requested builds may still be in-progress, or may have finished\nsuccessfully or unsuccessfully.",
	//   "flatPath": "v1/projects/{projectId}/builds",
	//   "httpMethod": "GET",
	//   "id": "cloudbuild.projects.builds.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The raw filter text to constrain the results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Number of results to return in the list.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to provide to skip to a particular spot in the list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/builds",
	//   "response": {
	//     "$ref": "ListBuildsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsBuildsListCall) Pages(ctx context.Context, f func(*ListBuildsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudbuild.projects.builds.retry":

type ProjectsBuildsRetryCall struct {
	s                 *Service
	projectId         string
	id                string
	retrybuildrequest *RetryBuildRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Retry: Creates a new build based on the specified build.
//
// This method creates a new build using the original build request,
// which may
// or may not result in an identical build.
//
// For triggered builds:
//
// * Triggered builds resolve to a precise revision; therefore a retry
// of a
// triggered build will result in a build that uses the same
// revision.
//
// For non-triggered builds that specify `RepoSource`:
//
// * If the original build built from the tip of a branch, the retried
// build
// will build from the tip of that branch, which may not be the same
// revision
// as the original build.
// * If the original build specified a commit sha or revision ID, the
// retried
// build will use the identical source.
//
// For builds that specify `StorageSource`:
//
// * If the original build pulled source from Google Cloud Storage
// without
// specifying the generation of the object, the new build will use the
// current
// object, which may be different from the original build source.
// * If the original build pulled source from Cloud Storage and
// specified the
// generation of the object, the new build will attempt to use the
// same
// object, which may or may not be available depending on the
// bucket's
// lifecycle management settings.
func (r *ProjectsBuildsService) Retry(projectId string, id string, retrybuildrequest *RetryBuildRequest) *ProjectsBuildsRetryCall {
	c := &ProjectsBuildsRetryCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.id = id
	c.retrybuildrequest = retrybuildrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBuildsRetryCall) Fields(s ...googleapi.Field) *ProjectsBuildsRetryCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBuildsRetryCall) Context(ctx context.Context) *ProjectsBuildsRetryCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsBuildsRetryCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsBuildsRetryCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.retrybuildrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/builds/{id}:retry")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"id":        c.id,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.builds.retry" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsBuildsRetryCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new build based on the specified build.\n\nThis method creates a new build using the original build request, which may\nor may not result in an identical build.\n\nFor triggered builds:\n\n* Triggered builds resolve to a precise revision; therefore a retry of a\ntriggered build will result in a build that uses the same revision.\n\nFor non-triggered builds that specify `RepoSource`:\n\n* If the original build built from the tip of a branch, the retried build\nwill build from the tip of that branch, which may not be the same revision\nas the original build.\n* If the original build specified a commit sha or revision ID, the retried\nbuild will use the identical source.\n\nFor builds that specify `StorageSource`:\n\n* If the original build pulled source from Google Cloud Storage without\nspecifying the generation of the object, the new build will use the current\nobject, which may be different from the original build source.\n* If the original build pulled source from Cloud Storage and specified the\ngeneration of the object, the new build will attempt to use the same\nobject, which may or may not be available depending on the bucket's\nlifecycle management settings.",
	//   "flatPath": "v1/projects/{projectId}/builds/{id}:retry",
	//   "httpMethod": "POST",
	//   "id": "cloudbuild.projects.builds.retry",
	//   "parameterOrder": [
	//     "projectId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Build ID of the original build.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/builds/{id}:retry",
	//   "request": {
	//     "$ref": "RetryBuildRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.triggers.create":

type ProjectsTriggersCreateCall struct {
	s            *Service
	projectId    string
	buildtrigger *BuildTrigger
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Create: Creates a new `BuildTrigger`.
//
// This API is experimental.
func (r *ProjectsTriggersService) Create(projectId string, buildtrigger *BuildTrigger) *ProjectsTriggersCreateCall {
	c := &ProjectsTriggersCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.buildtrigger = buildtrigger
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTriggersCreateCall) Fields(s ...googleapi.Field) *ProjectsTriggersCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTriggersCreateCall) Context(ctx context.Context) *ProjectsTriggersCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTriggersCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTriggersCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.buildtrigger)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/triggers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.triggers.create" call.
// Exactly one of *BuildTrigger or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *BuildTrigger.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTriggersCreateCall) Do(opts ...googleapi.CallOption) (*BuildTrigger, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BuildTrigger{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new `BuildTrigger`.\n\nThis API is experimental.",
	//   "flatPath": "v1/projects/{projectId}/triggers",
	//   "httpMethod": "POST",
	//   "id": "cloudbuild.projects.triggers.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "ID of the project for which to configure automatic builds.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/triggers",
	//   "request": {
	//     "$ref": "BuildTrigger"
	//   },
	//   "response": {
	//     "$ref": "BuildTrigger"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.triggers.delete":

type ProjectsTriggersDeleteCall struct {
	s          *Service
	projectId  string
	triggerId  string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a `BuildTrigger` by its project ID and trigger
// ID.
//
// This API is experimental.
func (r *ProjectsTriggersService) Delete(projectId string, triggerId string) *ProjectsTriggersDeleteCall {
	c := &ProjectsTriggersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.triggerId = triggerId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTriggersDeleteCall) Fields(s ...googleapi.Field) *ProjectsTriggersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTriggersDeleteCall) Context(ctx context.Context) *ProjectsTriggersDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTriggersDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTriggersDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/triggers/{triggerId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"triggerId": c.triggerId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.triggers.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsTriggersDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a `BuildTrigger` by its project ID and trigger ID.\n\nThis API is experimental.",
	//   "flatPath": "v1/projects/{projectId}/triggers/{triggerId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudbuild.projects.triggers.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "triggerId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "ID of the project that owns the trigger.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "ID of the `BuildTrigger` to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/triggers/{triggerId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.triggers.get":

type ProjectsTriggersGetCall struct {
	s            *Service
	projectId    string
	triggerId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns information about a `BuildTrigger`.
//
// This API is experimental.
func (r *ProjectsTriggersService) Get(projectId string, triggerId string) *ProjectsTriggersGetCall {
	c := &ProjectsTriggersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.triggerId = triggerId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTriggersGetCall) Fields(s ...googleapi.Field) *ProjectsTriggersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTriggersGetCall) IfNoneMatch(entityTag string) *ProjectsTriggersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTriggersGetCall) Context(ctx context.Context) *ProjectsTriggersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTriggersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTriggersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/triggers/{triggerId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"triggerId": c.triggerId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.triggers.get" call.
// Exactly one of *BuildTrigger or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *BuildTrigger.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTriggersGetCall) Do(opts ...googleapi.CallOption) (*BuildTrigger, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BuildTrigger{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns information about a `BuildTrigger`.\n\nThis API is experimental.",
	//   "flatPath": "v1/projects/{projectId}/triggers/{triggerId}",
	//   "httpMethod": "GET",
	//   "id": "cloudbuild.projects.triggers.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "triggerId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "ID of the project that owns the trigger.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "ID of the `BuildTrigger` to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/triggers/{triggerId}",
	//   "response": {
	//     "$ref": "BuildTrigger"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.triggers.list":

type ProjectsTriggersListCall struct {
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists existing `BuildTrigger`s.
//
// This API is experimental.
func (r *ProjectsTriggersService) List(projectId string) *ProjectsTriggersListCall {
	c := &ProjectsTriggersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTriggersListCall) Fields(s ...googleapi.Field) *ProjectsTriggersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTriggersListCall) IfNoneMatch(entityTag string) *ProjectsTriggersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTriggersListCall) Context(ctx context.Context) *ProjectsTriggersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTriggersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTriggersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/triggers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.triggers.list" call.
// Exactly one of *ListBuildTriggersResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListBuildTriggersResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsTriggersListCall) Do(opts ...googleapi.CallOption) (*ListBuildTriggersResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListBuildTriggersResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists existing `BuildTrigger`s.\n\nThis API is experimental.",
	//   "flatPath": "v1/projects/{projectId}/triggers",
	//   "httpMethod": "GET",
	//   "id": "cloudbuild.projects.triggers.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "ID of the project for which to list BuildTriggers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/triggers",
	//   "response": {
	//     "$ref": "ListBuildTriggersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.triggers.patch":

type ProjectsTriggersPatchCall struct {
	s            *Service
	projectId    string
	triggerId    string
	buildtrigger *BuildTrigger
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Patch: Updates a `BuildTrigger` by its project ID and trigger
// ID.
//
// This API is experimental.
func (r *ProjectsTriggersService) Patch(projectId string, triggerId string, buildtrigger *BuildTrigger) *ProjectsTriggersPatchCall {
	c := &ProjectsTriggersPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.triggerId = triggerId
	c.buildtrigger = buildtrigger
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTriggersPatchCall) Fields(s ...googleapi.Field) *ProjectsTriggersPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTriggersPatchCall) Context(ctx context.Context) *ProjectsTriggersPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTriggersPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTriggersPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.buildtrigger)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/triggers/{triggerId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"triggerId": c.triggerId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.triggers.patch" call.
// Exactly one of *BuildTrigger or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *BuildTrigger.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTriggersPatchCall) Do(opts ...googleapi.CallOption) (*BuildTrigger, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BuildTrigger{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a `BuildTrigger` by its project ID and trigger ID.\n\nThis API is experimental.",
	//   "flatPath": "v1/projects/{projectId}/triggers/{triggerId}",
	//   "httpMethod": "PATCH",
	//   "id": "cloudbuild.projects.triggers.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "triggerId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "ID of the project that owns the trigger.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "ID of the `BuildTrigger` to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/triggers/{triggerId}",
	//   "request": {
	//     "$ref": "BuildTrigger"
	//   },
	//   "response": {
	//     "$ref": "BuildTrigger"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.triggers.run":

type ProjectsTriggersRunCall struct {
	s          *Service
	projectId  string
	triggerId  string
	reposource *RepoSource
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Run: Runs a `BuildTrigger` at a particular source revision.
func (r *ProjectsTriggersService) Run(projectId string, triggerId string, reposource *RepoSource) *ProjectsTriggersRunCall {
	c := &ProjectsTriggersRunCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.triggerId = triggerId
	c.reposource = reposource
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTriggersRunCall) Fields(s ...googleapi.Field) *ProjectsTriggersRunCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTriggersRunCall) Context(ctx context.Context) *ProjectsTriggersRunCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTriggersRunCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTriggersRunCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.reposource)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/triggers/{triggerId}:run")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"triggerId": c.triggerId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudbuild.projects.triggers.run" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTriggersRunCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Runs a `BuildTrigger` at a particular source revision.",
	//   "flatPath": "v1/projects/{projectId}/triggers/{triggerId}:run",
	//   "httpMethod": "POST",
	//   "id": "cloudbuild.projects.triggers.run",
	//   "parameterOrder": [
	//     "projectId",
	//     "triggerId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "ID of the trigger.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/triggers/{triggerId}:run",
	//   "request": {
	//     "$ref": "RepoSource"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
