import pytest

from compare_pdfs import compare_pdf_files, get_version

SAMPLE_FILE_1_EXPECTED_OUTPUT = {
    "files": [
        {
            "filename": "sample_file_1.pdf",
            "path_to_file": "./sample_files",
            "n_pages": 1,
            "n_suspicious_pages": 1,
            "suspicious_pages": [1],
        },
        {
            "filename": "sample_file_1.pdf",
            "path_to_file": "./sample_files",
            "n_pages": 1,
            "n_suspicious_pages": 1,
            "suspicious_pages": [1],
        },
    ],
    "suspicious_pairs": [
        {
            "type": "Duplicate page",
            "cosine_distance": 0.0,
            "page_text": "* Critical Care: This is a specific referral to board-certified critical care specialist. If it is after hours, there will be a call-in fee. ** Please note cardiology availability changes from week to week and may not be available every week. Please contact us about availability. ***Please call 604-514-8383 when sending direct transfers",
            "pages": [
                {
                    "file_index": 0,
                    "page": 1,
                    "bbox": (0.01, 0.01, 0.99, 0.99),
                },
                {
                    "file_index": 1,
                    "page": 1,
                    "bbox": (0.01, 0.01, 0.99, 0.99),
                },
            ],
            "importance": 0,
        }
    ],
    "num_suspicious_pairs": 1,
    "similarity_scores": {
        0: {
            1: {
                "Common digit sequence": 0.0,
                "Common text string": 0.0,
                "Identical image": "Undefined",
            }
        }
    },
    "version": get_version(),
}


@pytest.mark.parametrize(
    "filenames, expected_output",
    [
        (
            ["./sample_files/sample_file_1.pdf", "./sample_files/sample_file_1.pdf"],
            SAMPLE_FILE_1_EXPECTED_OUTPUT,
        )
    ],
)
def test_compare_pdf_files(filenames, expected_output):
    """Test Compare pdf files"""
    actual_output = compare_pdf_files(filenames, regen_cache=True, pretty_print=True)
    keys_to_pop = ["time", "elapsed_time_sec", "pages_per_second"]
    for key in keys_to_pop:
        actual_output.pop(key, None)
    assert actual_output == expected_output


def test_get_version():
    """Test get version"""
    actual_output = get_version()
    assert actual_output != ""
