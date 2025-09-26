class BoutiqueInventory
  require 'ostruct'
  attr_reader :items

  def initialize(items)
    @items = items
    @items.map! {|item| OpenStruct.new(item)}
  end

  def item_names
    items.map { |item| item.name }.sort
  end

  def total_stock
    items.map{|item| item.quantity_by_size.values }.flatten.sum
  end
end
