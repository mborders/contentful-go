v0.5.0 (unreleased)
===
(TODO: changes clean-up)

* Added uploadFrom property to the asset resource.
* Made Locale field of assets exportable.
* Added archive properties to the Sys struct
* Made locale field in Entry model exportable
* Added missing X-Contentful-Content-Type header.
* updated gomod
* Added extra unit test.
* Added support for deleting resources.
* Added support for creating/uploading resources.
* Added support for getting an uploaded resource.
* [Syntax], [CP-182] Renamed some files, added unit tests to the Get method of ContentTypes service.
* [Syntax] Fixed some warnings and syntax notifications.
* [CP-182] Added a second unit test for get methods.
* [CP-182] Coverage improvements.
* [CP-182] Coverage improvements.
* [CP-182] Added missing unit tests to the Access Token and User services.
* [CP-182] Added missing unit tests to the Access Token and User services.
* [CP-182] Global coverage improvements.
* [CP-182] Coverage improvements for the collection service.
* [CP-182] Coverage improvements for the webhook service.
* [CP-182] Added more unit tests, and improved coverage.
* [CP-182] Added unit tests for the upsert method of the asset service.
* [CP-182] Added unit tests for the upsert method of the asset service.
* [[Syntax] Syntax improvements
* [[Syntax] Renamed wrong comments in app definition service.
* [CP-182] Adding missing asserts in the APIKey service.
* [CP-182] Adding missing unit tests of the upsert method in APIKeys service.
* [CP-77] Added support for getting usage statistics from an organization by spaces.
* [CP-76] Added support for getting usage statistics from an organization.
* [CP-75] Added support for deleting app installations.
* [CP-73] Added support for creating and updating app installations.
* [CP-72], [CP-74] Added support for listing and getting app installations from the Content Management API.
* [CP-71] Added support for deleting app definitions.
* Create method unit test asserts added.
* [CP-68], [CP-70] Added support for creating and updating app definitions.
* [CP-69] Added support for getting individual app definitions.
* [CP-67] Added support for listing app definitions.
* [CP-39], [CP-40] Added support for getting call details and health.
* [CP-38] Added support for listing an overview of all recent calls to a webhook.
* [CP-26] Added support for deleting extensions.
* [CP-23], [CP-24] Added support for creating and updating extensions.
* [CP-22] Added support for getting single extensions.
* [CP-22] Added support for listing all extensions.
* [CP-21] Added support for updating editor interfaces.
* [CP-20] Added support for getting single editor interfaces.
* [CP-19] Added support for listing editor interfaces.
* [CP-65], [CP-66] Added support for creating and deleting scheduled actions of entries.
* [CP-64] Added support for getting all scheduled actions from CMA.
* [CP-60], [CP-62], [CP-63] Added support for creating, updating and deleting entry tasks.
* [CP-61] Added support for getting a single entry task from the CMA.
* [CP-59] Added support for listing entry tasks.
* [Syntax] Cleared most warnings inside the codebase.
* [CP-57] Added support for revoking access tokens.
* [CP-53], [CP-54] Added missing unit tests.
* [CP-57] Added support for creating access tokens.
* [CP-55], [CP-56] Added support for getting and listing access tokens.
* [CP-46], [CP-47], [CP-48], [CP-49] Added support for getting and listing snapshots.
* [CP-36], [CP-37] Added support for archiving and unarchiving assets.
* [CP-34], [CP-35] Added support for listing only published assets and unpublishing assets. Also wrote all unit tests for the assets service.
* [CP-52] Added support for deleting space memberships.
* [CP-51] Added support for creating and updating memberships.
* [CP-50] Added support for getting a single space membership.
* [CP-45] Added support for listing space memberships.
* [CP-45] Added support for deleting a role of a space.
* [CP-44] Added support for updating a role of a space.
* [CP-42] Added support for creating a new role of a space.
* [CP-43] Added support for getting a single role of a space.
* [CP-41] Added support for listing all roles of one space.
* [CP-29], [CP-30] Added support for archiving and unarchiving entries.
* [CP-18], [CP-27], [CP-28] Added support for creating and updating entries in a space, also included is the functionality to get only all activated content types.
* [CP-17] Added getting all organizations of authenticated user.
* [CP-10] Added support for updating environment aliases.
* [CP-8], [CP-9] Added support for listing and getting environment aliases.
* [CP-7] Delete functionality of environments.
* [CP-4], [CP-6] Upsert functionality of environments.
* [CP-6] Get one single environment of space.
* [CP-1] Get all environments of space.
* [CP-58] Get the authenticated user.
* Update moul/http2curl dependency (#44)
* Update Gopkg.lock
* Remove old http2curl module from go.sum
* Switch from github.com/moul/http2curl to moul.io/http2curl
* added environment + improved memory alloc (#39)
* removed checking for environment
* added environment + improved memory alloc


v0.4.0 (2019-04-29)
===
* `+` Added support for creating/uploading resources.
* `+` Adding test for contentful.SetClient
* `~` Allowing the HTTP client to be set by the consumer. Useful for testing or implementing more robust HTTP clients.
* `+` Add Go Modules
* `+` added field type Symbol
* `+` Created an asset Alias for custom UnmarshalJSON method and fixed typo
* `+` Add MIT License
* `~` Capture validation errors
* `+` New content types can specify their ID.
* `~` Rename methods Activate/Deactivate to Publish/Unpublish
* `+` Extend EntryService by Delete, Publish and Unpublish actions


v0.3.1 (2017-11-28)
===
* `~` sdk version header format fixed.

v0.3.0 (2017-11-11)
===
* `~` use codecov as coverage service
* `+` `golint` is added to the CI process
* `~` `dep` is updated to the latest version
* `x` `vendor` folder is not under version control
* `~` `makefile` simplifications
* `+` testing and linting is now handled by scripts under `tools` folder
* `~` cosmetic changes in codebase to make linter happy


v0.2.0 (2017-04-12)
===
* Godoc style examples
* [Added] Query.go tests
* [Added] Locale resource tests
* [Added] ContentType tests
* [Added] Missing space resource tests
* [Added] User-Agent for api requests

v0.1.1 (2017-03-31)
===

* [Added] Rate-limited api requests
* [Fix] Locale model
* [Added] Content type field unmarshaling

v0.1.0 (2017-03-26)
===

### Introducing resource services
Every entity resource now has its own service definition for handling api communication. With this release, we don't store `contentful client` and `space` objects inside entities anymore. Resource services now get `spaceID` as a string parameter when it is neccessary.

With the old versions, in order to create a new `ContentType`, for example, you first need to observe `Space` object. That is no longer required. The problem with the old method was that you had to make an extra api request to observe the `Space` in order to interact with rest of the resources. The following example demonstras the difference between old and new version.

```go
// prior to v0.1.0
space, err := cma.GetSpace("space-id") // this call was making an extra api call
contentTypes, err := space.ContentTypes()

// after v0.1.0
contentType := &contentful.ContentType{ ... }
spaceID := "space-id"
cma.ContentTypes.Upsert(spaceID, contentType) // now we are passing spaceID as string
```

You can access available resources as follows:

```go
cma := contentful.NewCMA("token")
cma.Spaces
cma.APIKeys
cma.Assets
cma.ContentTypes
cma.Entries
cma.Locales
cma.Webhooks
```

Every resource service exposes the following interface:

`List(spaceID string) *Collection`

`Get(spaceID, contentTypeID string)(*ContentType, error)`

`Upsert(spaceID string, ct *ContentType) error`

`Delete(spaceID string, ct *ContentType) error`

### Create resource instancas directly from their model definitions

All `New{ResourceName}`, such as `NewContentType`, `NewSpace`, functions are removed from the SDK. It turned out that it wasn't a good practice in golang. Instead, you can directly initiate resource entities directly from their models, such as:

```go
contentType := &contentful.ContentType{
    Name: "name",
    ... other fields
}
```


v0.0.3 (2017-03-22)
===
* [Added] PredefinedValues validation
* [Added] Range validation greater/less than equal to support.
* [Added] Size validation for content type field.
* [Added] Packages are vendored with `godep`.
* [Added] `version.go`.
* [Added] `entity/content_type`: regex validation for content type field.
* [Added] Validation data structures added: `MinMax`, `Regex`
* [Added] `LinkType` support for `Field` struct
* [Added] New validations: `MimeType`, `Dimension`, `FileSize`


v0.0.2 (2017-03-21)
===
* `entity/webhook`: add tests for webhook entity.
* `entity/space`: add tests for space entity.
* `errors`: add tests for error handler.
* `entity/content_type`: add test for content type entity.
* `entity/content_type`: Field validations added for link type
* `entity/content_type`: field validations added: Range, PredefinedValues, Unique


v0.0.1 (2017-03-20)
===
* `sdk`: first implementation.
* `collection`: first implementation.
* `entity/content_type`: first implementation.
* `entity/entry`: first implementation.
* `entity/query`: first implementation.
* `entity/asset`: first implementation.
* `entity/locale`: first implementation.
* `entity/space`: first implementation.
* `entity/webhook`: first implementation.
* `entity/api_key`: first implementation.
* `sdk`: basic documentation.
* `examples`: some examples for entities
