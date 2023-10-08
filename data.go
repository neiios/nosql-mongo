package main

var hotels = []Hotel{
	{
		Name:    "Prestige Hotel",
		Address: "123 Luxury St.",
		Country: "USA",
		Rooms: []Room{
			{
				Price:  200,
				Booked: false,
			},
			{
				Price:  250,
				Booked: true,
			},
		},
		Workers: []Worker{
			{
				Name:     "John Doe",
				Age:      30,
				Position: "Manager",
			},
			{
				Name:     "Jane Smith",
				Age:      25,
				Position: "Receptionist",
			},
		},
	},
	{
		Name:    "Ocean View Resort",
		Address: "456 Beach Ave.",
		Country: "Bahamas",
		Rooms: []Room{
			{
				Price:  350,
				Booked: true,
			},
			{
				Price:  300,
				Booked: false,
			},
		},
		Workers: []Worker{
			{
				Name:     "Alice Johnson",
				Age:      40,
				Position: "Manager",
			},
			{
				Name:     "Bob Martin",
				Age:      22,
				Position: "Lifeguard",
			},
		},
	},
}
