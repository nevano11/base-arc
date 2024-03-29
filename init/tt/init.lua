box.cfg{}

box.schema.space.create('bands')

box.space.bands:format({
    { name = 'id', type = 'unsigned' },
    { name = 'band_name', type = 'string' },
    { name = 'year', type = 'unsigned' }
})

box.space.bands:create_index('primary', { type = "tree", parts = { 'id' } })

box.space.bands:insert { 1, 'Roxette', 1986 }
box.space.bands:insert { 2, 'Scorpions', 1965 }
box.space.bands:insert { 3, 'Ace of Base', 1987 }

box.space.bands:create_index('secondary', { type = "tree", parts = { 'band_name' } })