# line - nazwa linii
# departure_time - czas odjazdu w formacie HH:MM:SS
# arrival_time - czas przyjazdu w formacie HH:MM:SS
# start_stop - przystanek początkowy
# end_stop - przystanek końcowy
# start_stop_lat - szerokość geograficzna przystanku początkowego 
# start_stop_lon - długość geograficzna przystanku początkowego
# end_stop_lat - szerokość geograficzna przystanku końcowego
# end_stop_lon - długość geograficzna przystanku końcowego 

class Edge:
  def __init__(self, row: str) -> None:
    self.line, 
    self.departure_time, 
    self.arrival_time, 
    self.start_stop, 
    self.end_stop, 
    self.start_stop_lat, 
    self.start_stop_lon, 
    self.end_stop_lat, 
    self.end_stop_lon = row.split(',')

    self.departure_time_sec = self.time_to_seconds(self.departure_time)

  def __str__(self) -> str:
    return f'{self.line} {self.departure_time} {self.arrival_time} {self.start_stop} {self.end_stop} {self.start_stop_lat} {self.start_stop_lon} {self.end_stop_lat} {self.end_stop_lon}'
  
  def time_to_seconds(self, time: str) -> int:
    h, m, s = map(int, time.split(':'))
    return (h * 3600 + m * 60 + s) % 86400
  

