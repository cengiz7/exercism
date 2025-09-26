class LocomotiveEngineer
  def self.generate_list_of_wagons(*wagon_ids)
    wagon_ids
  end

  def self.fix_list_of_wagons(each_wagons_id, missing_wagons)
    first, second, third, *rest = each_wagons_id
    result = third, *missing_wagons, *rest, first, second
  end

  def self.add_missing_stops(routing, **stops)

    stops_list =
      stops.each_pair
        .sort { | key, _ | key.to_s.gsub('stop_', '').to_i }
        .map { | _, value | value }

    routing.merge({ stops: stops_list })
  end

  def self.extend_route_information(route, more_route_information)
    {**route, **more_route_information}
  end
end
