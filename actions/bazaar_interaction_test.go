package actions

import (
	"math"
	"strconv"
	"testing"

	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/errors"

	"github.com/google/go-cmp/cmp"
)

type FakeItemGateway struct {
}

func (i FakeItemGateway) SearchByVNum(vnum int) (entities.Item, error) {
	if vnum != 5 {
		return entities.Item{}, errors.ItemNotFound{VNum: vnum}
	}

	return entities.Item{VNum: 5, InventoryPocket: 1}, nil
}

type FakeBazaarGateway struct {
	isOpen bool
}

func (b *FakeBazaarGateway) Open(int) error {
	b.isOpen = true

	return nil
}

func (b *FakeBazaarGateway) Close() error {
	b.isOpen = false

	return nil
}

func (b FakeBazaarGateway) IsOpen() bool {
	return b.isOpen
}

func (b FakeBazaarGateway) SearchItemsByVNumAndPage(vnums []int, _ int) ([]entities.BazaarItem, error) {
	for i := range vnums {
		if vnums[i] != 5 {
			return nil, errors.ItemNotFound{VNum: vnums[i]}
		}
	}

	return []entities.BazaarItem{{
		ItemInstance: entities.ItemInstance{
			Item: entities.Item{
				VNum:            5,
				InventoryPocket: 1,
			},
			Amount:    1,
			OwnerID:   567,
			OwnerName: "test",
		},
		Price:       100,
		MinutesLeft: 5,
		SaleStatus:  enums.All,
		SoldAmount:  0,
	}}, nil
}

type FakeCharacterGateway struct {
	character entities.Character
}

func (c FakeCharacterGateway) Info() entities.Character {
	return c.character
}

func (c FakeCharacterGateway) WalkTo(entities.Point) error {
	return nil
}

func (c FakeCharacterGateway) CanMove() bool {
	return true
}

func (c FakeCharacterGateway) CanAttack() bool {
	return true
}

type FakeMapGateway struct {
}

func (m FakeMapGateway) Info() entities.Map {
	return entities.Map{
		ID:     0,
		Width:  math.MaxInt32,
		Height: math.MaxInt32,
	}
}

func (m FakeMapGateway) DistanceBetween(p1 entities.Point, p2 entities.Point) (int, error) {
	if p1.X == p2.X && p1.Y == p2.Y {
		return 0, nil
	} else if p1.X < 0 || p1.Y < 0 || p2.X < 0 || p2.Y < 0 {
		return -1, errors.NoPathToPoint{From: p1, To: p2}
	}

	return math.MaxInt32, nil
}

func (m FakeMapGateway) FindPath(entities.Point, entities.Point) ([]entities.Point, error) {
	return []entities.Point{}, nil
}

func (m FakeMapGateway) IsWalkable(entities.Point) bool {
	return true
}

type FakeNPCGateway struct {
	bazaarNPC entities.NPC
}

func (n FakeNPCGateway) All() []entities.NPC {
	return []entities.NPC{n.bazaarNPC}
}

func (n FakeNPCGateway) SearchByID(id int) (entities.NPC, bool) {
	if id != n.bazaarNPC.ID {
		return entities.NPC{}, false
	}

	return n.bazaarNPC, true
}

type FakeShopGateway struct {
	bazaarShop entities.Shop
}

func (s FakeShopGateway) All() []entities.Shop {
	return []entities.Shop{s.bazaarShop}
}

func (s FakeShopGateway) SearchByID(id int) (entities.Shop, bool) {
	if id != s.bazaarShop.OwnerID {
		return entities.Shop{}, false
	}

	return s.bazaarShop, true
}

func (s FakeShopGateway) SearchByShopType(st enums.ShopType) (entities.Shop, bool) {
	if st != s.bazaarShop.ShopType {
		return entities.Shop{}, false
	}

	return s.bazaarShop, true
}

var bazaarOpenCloseTests = []struct {
	character  entities.Character
	bazaarNPC  entities.NPC
	bazaarShop entities.Shop
	shouldWork bool
}{
	{
		character: entities.Character{},
		bazaarNPC: entities.NPC{
			Actor: entities.Actor{
				Thing: entities.Thing{
					ID: 1234,
				},
			},
		},
		bazaarShop: entities.Shop{
			OwnerID:  1234,
			Name:     "nosbazaar",
			ShopType: enums.NosBazaar,
		},
		shouldWork: true,
	},
	{
		character:  entities.Character{},
		bazaarNPC:  entities.NPC{},
		bazaarShop: entities.Shop{},
		shouldWork: false,
	},
	{
		character: entities.Character{},
		bazaarNPC: entities.NPC{
			Actor: entities.Actor{
				Thing: entities.Thing{
					Position: entities.Point{
						X: -10,
						Y: -10,
					},
					ID: 1234,
				},
			},
		},
		bazaarShop: entities.Shop{
			OwnerID:  1234,
			Name:     "nosbazaar",
			ShopType: enums.NosBazaar,
		},
		shouldWork: false,
	},
	{
		character: entities.Character{
			Player: entities.Player{
				Actor: entities.Actor{
					Thing: entities.Thing{
						Position: entities.Point{
							X: 100,
							Y: 100,
						},
					},
				},
			},
		},
		bazaarNPC: entities.NPC{
			Actor: entities.Actor{
				Thing: entities.Thing{
					ID: 1234,
				},
			},
		},
		bazaarShop: entities.Shop{
			OwnerID:  1234,
			Name:     "nosbazaar",
			ShopType: enums.NosBazaar,
		},
		shouldWork: false,
	},
}

func TestBazaarOpen(t *testing.T) {
	for i, tt := range bazaarOpenCloseTests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NewBazaarInteractor(
				FakeItemGateway{},
				&FakeBazaarGateway{},
				FakeCharacterGateway{character: tt.character},
				FakeMapGateway{},
				FakeNPCGateway{bazaarNPC: tt.bazaarNPC},
				FakeShopGateway{bazaarShop: tt.bazaarShop},
			)
			err := b.Open()
			if tt.shouldWork && err != nil {
				t.Errorf("%s", err.Error())
			} else if !tt.shouldWork && err == nil {
				t.Errorf("shouldn't have worked, but it did")
			}
		})
	}
}

func TestBazaarClose(t *testing.T) {
	for i, tt := range bazaarOpenCloseTests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := NewBazaarInteractor(
				FakeItemGateway{},
				&FakeBazaarGateway{},
				FakeCharacterGateway{character: tt.character},
				FakeMapGateway{},
				FakeNPCGateway{bazaarNPC: tt.bazaarNPC},
				FakeShopGateway{bazaarShop: tt.bazaarShop},
			)
			_ = b.Open()
			err := b.Close()
			if err != nil {
				t.Errorf("%s", err.Error())
			}
		})
	}
}

var bazaarSearchByVNumTests = []struct {
	vnum       int
	shouldWork bool
}{
	{
		vnum:       5,
		shouldWork: true,
	},
	{
		vnum:       1,
		shouldWork: false,
	},
}

func TestBazaarSearchByVNum(t *testing.T) {
	items := []entities.BazaarItem{{
		ItemInstance: entities.ItemInstance{
			Item: entities.Item{
				VNum:            5,
				InventoryPocket: 1,
			},
			Amount:    1,
			OwnerID:   567,
			OwnerName: "test",
		},
		Price:       100,
		MinutesLeft: 5,
		SaleStatus:  enums.All,
		SoldAmount:  0,
	}}

	b := NewBazaarInteractor(
		FakeItemGateway{},
		&FakeBazaarGateway{},
		FakeCharacterGateway{character: entities.Character{}},
		FakeMapGateway{},
		FakeNPCGateway{bazaarNPC: entities.NPC{}},
		FakeShopGateway{bazaarShop: entities.Shop{ShopType: enums.NosBazaar}},
	)

	err := b.Open()
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	for i, tt := range bazaarSearchByVNumTests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res, err := b.SearchItemByVNum(tt.vnum)
			if tt.shouldWork && err != nil {
				t.Errorf("%s", err.Error())
			} else if !tt.shouldWork && err == nil {
				t.Errorf("should've failed, but it worked")
			} else if tt.shouldWork && !cmp.Equal(items, res) {
				t.Errorf("expected: %v\nfound: %v\n", items, res)
			}
		})
	}
}
