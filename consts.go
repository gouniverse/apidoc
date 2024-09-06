package apidoc

// Content types
const CONTENT_TYPE_JSON = "application/json"
const CONTENT_TYPE_FORM = "application/x-www-form-urlencoded"
const CONTENT_TYPE_MULTIPART = "multipart/form-data"
const CONTENT_TYPE_TEXT = "text/plain"

// Block Types
const BLOCK_TYPE_GROUP = "group"
const BLOCK_TYPE_ENDPOINT = "endpoint"
const BLOCK_TYPE_EXAMPLE = "example"
const BLOCK_TYPE_PARAMETER = "parameter"
const BLOCK_TYPE_RESPONSE = "response"
const BLOCK_TYPE_SECTION = "section"

// Locations

// PARAMETER_LOCATION_PATH represents a parameter located
// in the URL path (e.g. /users/:id).
const PARAMETER_LOCATION_PATH = "path"

// PARAMETER_LOCATION_QUERY represents a parameter located
// in the URL query string (e.g. ?id=1).
const PARAMETER_LOCATION_QUERY = "query"

// PARAMETER_LOCATION_HEADER represents a parameter located
// in the HTTP request header (e.g. Authorization).
const PARAMETER_LOCATION_HEADER = "header"

// PARAMETER_LOCATION_COOKIE represents a parameter located
// in the HTTP request cookie (e.g. token).
const PARAMETER_LOCATION_COOKIE = "cookie"

// PARAMETER_LOCATION_FORM represents a parameter located
// in the HTTP request body as form data (e.g. name=John).
const PARAMETER_LOCATION_FORM = "form"

// PARAMETER_LOCATION_BODY represents a parameter located
// in the HTTP request body as a JSON payload or other
// structured data (e.g. {"name": "John"}).
const PARAMETER_LOCATION_BODY = "body"

// Light themes
const THEME_DEFAULT = "bootstrap"
const THEME_CERULEAN = "cerulean"
const THEME_COSMO = "cosmo"
const THEME_FLATLY = "flatly"
const THEME_JOURNAL = "journal"
const THEME_LITERA = "litera"
const THEME_LUMEN = "lumen"
const THEME_LUX = "lux"
const THEME_MATERIA = "materia"
const THEME_MINTY = "minty"
const THEME_MORPH = "morph"
const THEME_PULSE = "pulse"
const THEME_QUARTZ = "quartz"
const THEME_SANDSTONE = "sandstone"
const THEME_SIMPLEX = "simplex"
const THEME_SKETCHY = "sketchy"
const THEME_SPACELAB = "spacelab"
const THEME_UNITED = "united"
const THEME_YETI = "yeti"
const THEME_ZEPHYR = "zephyr"

// Dark themes
const THEME_CYBORG = "cyborg"
const THEME_DARKLY = "darkly"
const THEME_SLATE = "slate"
const THEME_SOLAR = "solar"
const THEME_SUPERHERO = "superhero"
const THEME_VAPOR = "vapor"
