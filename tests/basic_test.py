from src.compare_pdfs_util import get_models_directory


def test_get_models_directory():
    path = get_models_directory()
    assert path != ""
