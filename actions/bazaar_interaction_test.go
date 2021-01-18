package actions_test

import (
	"strconv"
	"testing"

	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/entities"
	"github.com/gilgames000/go-noskit/enums"
	"github.com/gilgames000/go-noskit/testdoubles"

	"github.com/google/go-cmp/cmp"
)

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
			b := actions.NewBazaarInteractor(
				testdoubles.FakeItemGateway{},
				&testdoubles.FakeBazaarGateway{},
				testdoubles.FakeCharacterGateway{Character: tt.character},
				testdoubles.FakeMapGateway{},
				testdoubles.FakeNPCGateway{BazaarNPC: tt.bazaarNPC},
				testdoubles.FakeShopGateway{BazaarShop: tt.bazaarShop},
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
			b := actions.NewBazaarInteractor(
				testdoubles.FakeItemGateway{},
				&testdoubles.FakeBazaarGateway{},
				testdoubles.FakeCharacterGateway{Character: tt.character},
				testdoubles.FakeMapGateway{},
				testdoubles.FakeNPCGateway{BazaarNPC: tt.bazaarNPC},
				testdoubles.FakeShopGateway{BazaarShop: tt.bazaarShop},
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

	b := actions.NewBazaarInteractor(
		testdoubles.FakeItemGateway{},
		&testdoubles.FakeBazaarGateway{},
		testdoubles.FakeCharacterGateway{Character: entities.Character{}},
		testdoubles.FakeMapGateway{},
		testdoubles.FakeNPCGateway{BazaarNPC: entities.NPC{}},
		testdoubles.FakeShopGateway{BazaarShop: entities.Shop{ShopType: enums.NosBazaar}},
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
