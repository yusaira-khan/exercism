class ResistorColorDuo
  def self.colorDashes
    "Black - Brown - Red - Orange - Yellow - Green - Blue - Violet - Grey - White"
    #0-9
  end
  def self.colorStr
    self.colorDashes.gsub("- ","").downcase
  end
  def self.colorList
    self.colorStr.split(" ")
  end
  def self.colorDict
    Hash[self.colorList.collect.with_index {|v,i| [v,i]}]
  end

  def self.value(colors)
    if colors.length < 2
      0
    else
      self.colorDict[colors[0]] * 10 + self.colorDict[colors[1]]
    end
  end

end
