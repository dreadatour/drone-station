export type Drone = {
  id: string
  quadrant: string
  x: string
  y: string
};

export type DroneAdd = {
  quadrant: string
  x: string
  y: string
};

export type DronesMap = Map<string, Drone>;
