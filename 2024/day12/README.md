Initial idea of just adding up perimeter based on neighbors being different works for case of a single region per crop, but multiple crops per region doesn't work.

Next thought:
* Walk each node of the garden map
  * If the node hasn't been walked yet, Walk a region from that starting point
  * Append each region to a slice instead of a map of Crops
