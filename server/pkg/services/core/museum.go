package core

// Interface to allow multiple museum packages to return a random artwork
type MuseumClient interface {
	GetRandomArtwork() (*Artwork, error)
}
