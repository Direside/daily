import re

def extract_title(line):
    return re.search(r"# (.+?)", line).group(1)

def main():
    title = ""
    yesterday = []
    today = []
    blocked = []

    with open("standup.md", "r") as s:
        for line in s:
            extract_title(line)

if __name__ == "__main__":
    main()
