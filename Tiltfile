
# Custom build tool for nix flakes
def build_flake_image(ref, path = "", output = "", resultfile = "result", deps = []):
    build_cmd = "nix build {path}#{output} --refresh --no-link --print-out-paths".format(
        path = path,
        output = output
    )
    commands = [
        "RESULT_IMAGE=$({cmd})".format(cmd = build_cmd),
        "docker image load -i ${RESULT_IMAGE}",
        'IMG_NAME="$(tar -Oxf $RESULT_IMAGE manifest.json | jq -r ".[0].RepoTags[0]")"'.format(ref = ref),
        "docker tag ${IMG_NAME} ${EXPECTED_REF}"
    ]
    custom_build(
        ref,
        command = [
            "nix-shell",
            "--packages",
            "coreutils",
            "gnutar",
            "jq",
            "--run",
            ";\n".join(commands),
        ],
        deps = deps,
    )


image_name = "kurtosistech/kardinal-manager"

build_flake_image(image_name , ".", "kardinal-manager-container", deps=["./kardinal-manager"])

yaml_dir = "./kardinal-manager/deployment"
k8s_yaml(yaml=(yaml_dir + "/k8s.yaml"))

if k8s_context:
    k8s_context(k8s_context)


