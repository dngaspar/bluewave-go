from src.utils import get_models_directory


def test_get_models_directory():
    path = get_models_directory()
    assert path != ""
