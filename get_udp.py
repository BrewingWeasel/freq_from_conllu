import requests
import argparse

parser = argparse.ArgumentParser("get_udp")
parser.add_argument("repo", help="The title of a universal dependencies repo", type=str)
parser.add_argument("language", help="The langauge of the repo", type=str)
args = parser.parse_args()

print("Getting UDP data")

json = requests.get(f"https://api.github.com/repos/UniversalDependencies/{args.repo}/git/trees/master?recursive=true").json()

with open(f"sources/{args.language}.conllu", "w") as f:
    for file in json["tree"]:
        if file["path"].endswith(".conllu"):
            print(f"Found UDP file: {file['path']}")
            contents = requests.get(f"https://raw.githubusercontent.com/UniversalDependencies/{args.repo}/master/{file['path']}")
            f.write(contents.text)
            f.write("\n\n")
            print("wrote to file")
