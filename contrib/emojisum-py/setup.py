from setuptools import find_packages, setup


def get_requirements():
    with open("requirements.txt") as f:
        return f.read().splitlines()


setup(
    name="emojisum",
    version="1.0.0",
    packages=find_packages(where="src"),
    url="",
    license="",
    author="Tamir Bahar",
    author_email="",
    description="",
    package_dir={"": "src"},
    include_package_data=True,
    install_requires=get_requirements(),
    entry_points={
        "console_scripts": ["emojisum=emojisum:entry"],
    },
)
