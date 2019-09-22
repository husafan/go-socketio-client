load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "45409e6c4f748baa9e05f8f6ab6efaa05739aa064e3ab94e5a1a09849c51806a",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.18.7/rules_go-0.18.7.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.18.7/rules_go-0.18.7.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
)

BAZEL_BUILD_TOOLS_SHA = "f27d1753c8b3210d9e87cdc9c45bc2739ae2c2db"

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "cdaac537b56375f658179ee2f27813cac19542443f4722b6730d84e4125355e6",
    strip_prefix = "buildtools-" + BAZEL_BUILD_TOOLS_SHA,
    url = "https://github.com/bazelbuild/buildtools/archive/" + BAZEL_BUILD_TOOLS_SHA + ".zip",
)

ATLASSIAN_BAZEL_TOOLS_SHA = "f04c7c0ec6851ea378f90c74c855bcb87d9793b5"

http_archive(
    name = "com_github_atlassian_bazel_tools",
    sha256 = "281aba7b549daaf418302cb0a51e8358d6557a03185019d6b521d9a12744c957",
    strip_prefix = "bazel-tools-" + ATLASSIAN_BAZEL_TOOLS_SHA,
    urls = ["https://github.com/atlassian/bazel-tools/archive/" + ATLASSIAN_BAZEL_TOOLS_SHA + ".zip"],
)

git_repository(
    name = "io_bazel_rules_docker",
    commit = "5d35fee42a513ac83dc23c32ba1f0b029fd75e0f",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    shallow_since = "1559569106 -0700",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

load("@com_github_atlassian_bazel_tools//goimports:deps.bzl", "goimports_dependencies")

goimports_dependencies()

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)
load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    tag = "v0.0.9",
)

go_repository(
    name = "com_github_toqoz_gopwt",
    importpath = "github.com/ToQoz/gopwt",
    tag = "v2.1.0",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "d0b11bdaac8a",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "com_github_franela_goblin",
    commit = "ead4ad1d2727",
    importpath = "github.com/franela/goblin",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    tag = "v0.0.4",
)

go_repository(
    name = "com_github_k0kubun_pp",
    importpath = "github.com/k0kubun/pp",
    tag = "v3.0.1",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    tag = "v0.1.2",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    tag = "v1.1.0",
)

go_repository(
    name = "com_github_k0kubun_colorstring",
    commit = "9440f1994b88",
    importpath = "github.com/k0kubun/colorstring",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_sergi_go_diff",
    importpath = "github.com/sergi/go-diff",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    tag = "v0.1.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    tag = "v1.4.0",
)

go_repository(
    name = "in_gopkg_check_v1",
    commit = "20d25e280405",
    importpath = "gopkg.in/check.v1",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    tag = "v2.2.2",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    commit = "0766667cb4d1",
    importpath = "github.com/gopherjs/gopherjs",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    tag = "v4.20.0",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    commit = "b2de0cb4f26d",
    importpath = "github.com/smartystreets/assertions",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    commit = "505e41936337",
    importpath = "github.com/smartystreets/goconvey",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "c2843e01d9a2",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "d8887717615a",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    tag = "v0.3.0",
)

go_repository(
    name = "org_golang_x_tools",
    commit = "ab21143f2384",
    importpath = "golang.org/x/tools",
)
