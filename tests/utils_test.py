import pytest

from src.utils import get_models_directory, list_of_unique_dicts, get_digits


def test_get_models_directory():
    """Check if function is returning the path of models directory"""
    path = get_models_directory()
    assert path != ""


@pytest.mark.parametrize(
    "data, expected_output",
    [
        ([{1: 2}, {1: 2}], [{1: 2}]),
        ([{1: 2, 3: 4}], [{1: 2, 3: 4}]),
        ([{1: 2, 3: 4}, {1: 2, 3: 4}], [{1: 2, 3: 4}]),
        ([{1: 2, 3: 4}, {1: 2, 3: 4}], [{3: 4, 1: 2}]),
        ([{"a": 1, "b": 2}, {1: 2, 3: 4}], [{"a": 1, "b": 2}, {3: 4, 1: 2}]),
    ],
)
def test_list_of_unique_dicts(data, expected_output):
    """Check if function is removing duplicates in a dict"""
    result = list_of_unique_dicts(data)
    assert result == expected_output


@pytest.mark.parametrize(
    "text, digits", [("abcd123", "123"), ("123", "123"), ("123.45", "12345")]
)
def test_get_digits(text, digits):
    """Test digits extraction from text"""
    result = get_digits(text)
    assert result == digits
