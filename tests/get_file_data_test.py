import pytest

from src.get_file_data import (
    page_skip_conditions,
    block_skip_conditions,
    is_compatible
)


@pytest.mark.parametrize(
    "text, expected_output",
    [
        ("", False),
        ("FORM FDA ", True),
        ("Safety Data Sheet", True),
        ("PAPERWORK REDUCTION ACT", True)
    ],
)
def test_page_skip_conditions(text, expected_output):
    """Test Page Skip Conditions"""
    actual_output = page_skip_conditions(text)
    assert any(actual_output) == expected_output


@pytest.mark.parametrize(
    "text, expected_output",
    [
        ("", True),
        ("510(k)", True),
        ("Fax:", True),
        ("phone", False)
    ],
)
def test_block_skip_conditions(text, expected_output):
    """Test Block Skip Conditions"""
    actual_output = block_skip_conditions(text)
    assert any(actual_output) == expected_output


@pytest.mark.parametrize(
    "text_1, text_2, expected_output",
    [
        ("1.1.1", "1.1.1", True),
        ("1.1.1", "1.2.1", False)
    ],
)
def test_is_compatible(text_1, text_2, expected_output):
    """Test is compatible"""
    actual_output = is_compatible(text_1, text_2)
    assert actual_output == expected_output
