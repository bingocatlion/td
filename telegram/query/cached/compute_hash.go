package cached

import (
	"sort"

	"github.com/gotd/td/telegram/query/hasher"
	"github.com/gotd/td/tg"
)

func (s *ContactsGetContacts) computeHash(v *tg.ContactsContacts) int64 {
	cts := v.Contacts

	sort.SliceStable(cts, func(i, j int) bool {
		return cts[i].UserID < cts[j].UserID
	})
	h := hasher.Hasher{}
	for _, contact := range cts {
		h.Update(uint32(contact.UserID))
	}

	return h.Sum()
}

func (s *MessagesGetQuickReplies) computeHash(v *tg.MessagesQuickReplies) int64 {
	r := v.QuickReplies

	sort.SliceStable(r, func(i, j int) bool {
		return r[i].ShortcutID < r[j].ShortcutID
	})
	h := hasher.Hasher{}
	for _, contact := range r {
		h.Update(uint32(contact.ShortcutID))
	}

	return h.Sum()
}

func (s *PaymentsGetStarGiftCollections) computeHash(v *tg.PaymentsStarGiftCollections) int64 {
	collections := v.Collections

	sort.SliceStable(collections, func(i, j int) bool {
		return collections[i].CollectionID < collections[j].CollectionID
	})

	h := hasher.Hasher{}
	for _, collection := range collections {
		h.Update(uint32(collection.CollectionID))
	}

	return h.Sum()
}

func (s *AccountGetSavedMusicIDs) computeHash(v *tg.AccountSavedMusicIDs) int64 {
	ids := v.IDs

	sort.SliceStable(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	h := hasher.Hasher{}
	for _, id := range ids {
		h.Update(uint32(id))
	}

	return h.Sum()
}
