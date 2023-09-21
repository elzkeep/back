package resources

type Tile struct {
	Type int
}

/*
class Tile {
	id: number;
	name: string;
	income: any;
	receive: any;
	action: any;
	use: boolean;

	constructor(id, name, income, receive, action) {
		this.id = id;
		this.name = name;
		this.income = income;
		this.receive = receive;
		this.action = action;
		this.use = false;
	}
}

		this.tiles.push(Tile(1, '4 coin', {coin: 4}, null, null));
		this.tiles.push(Tile(2, 'gaia point', null, {gaiaPoint: 3}, null));
		this.tiles.push(Tile(3, '1 coin 1 knowledge', {coin: 1, knowledge: 1}, null, null));
		this.tiles.push(Tile(4, '1 worker 1 power', {worker: 1, power: 1}, null, null));
		this.tiles.push(Tile(5, 'sh,sa power up', null, {buildingPower: [0, 0, 0, 1, 1, 1, 0]}, null));
		this.tiles.push(Tile(6, '4 power', null, null, {power: 4}));
		this.tiles.push(Tile(7, '1 worker 1 qic', null, {worker: 1, qic: 1}, null));
		this.tiles.push(Tile(8, '? knowledge', null, {knowledge: -1}, null));
		this.tiles.push(Tile(9, '7 point', null, {point: 7}, null));

		this.tiles.sort(function(){return 0.5-Math.random();});

*/
