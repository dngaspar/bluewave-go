# bluewave

This is a Python script to analyze the similarity of two PDFs.


## Installation

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install bluewave.

```bash
pip install -r requirements.txt
```

For installation of black in your local environment, please follow the instructions here:
```
https://black.readthedocs.io/en/stable/integrations/editors.html
```

## Usage
To run this, an example would be:

```bash
python compare_pdfs.py -f a.pdf b.pdf
```

```bash
python compare_pdfs.py -f a.pdf b.pdf -v
```

```bash
python compare_pdfs.py -f data/test_pdfs/00026_04_fda-K071597_test_data.pdf data/test_pdfs/small_test/copied_data.pdf
```

## Code Formatting
```bash
pre-commit run --all-files
```
