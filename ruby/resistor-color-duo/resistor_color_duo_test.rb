require 'minitest/autorun'
require_relative 'resistor_color_duo'
# Common test data version: 2.1.0 00dda3a

def assert_start_with full, sub
  assert (full.start_with? sub) , ("Expected #{full.inspect} to start with #{sub.inspect}")
end
class ResistorColorDuoTest < Minitest::Test
  def test_colorstr
    assert_start_with ResistorColorDuo.colorStr, "black brown"
  end
  def test_colorsplit
    assert_equal  "black", ResistorColorDuo.colorList[0]
  end
  def test_colordashes
    assert_start_with ResistorColorDuo.colorDashes, "Black - "
  end
  def test_colordict
    assert_equal   0, ResistorColorDuo.colorDict["black"]
    assert_equal   1, ResistorColorDuo.colorDict["brown"]
    assert_equal   6, ResistorColorDuo.colorDict["blue"]
    assert_equal   8, ResistorColorDuo.colorDict["grey"]
    assert_equal   4, ResistorColorDuo.colorDict["yellow"]
  end

  def test_color_empty
    assert_equal 0, ResistorColorDuo.value([])
  end

  def test_brown_and_black
    assert_equal 10, ResistorColorDuo.value(["brown", "black"])
  end

  def test_blue_and_grey
    assert_equal 68, ResistorColorDuo.value(["blue", "grey"])
  end

  def test_yellow_and_violet
    assert_equal 47, ResistorColorDuo.value(["yellow", "violet"])
  end

  def test_orange_and_orange
    assert_equal 33, ResistorColorDuo.value(["orange", "orange"])
  end

  def test_ignore_additional_colors
    assert_equal 51, ResistorColorDuo.value(["green", "brown", "orange"])
  end
end
