package streams_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/superseriousbusiness/activity/streams"
	vocab "github.com/superseriousbusiness/activity/streams/vocab"
)

var profileJSONIn = `{
  "@context": [
    "http://schema.org",
    "http://joinmastodon.org/ns",
    "https://w3id.org/security/v1",
    "https://www.w3.org/ns/activitystreams"
  ],
  "attachment": [
    {
      "name": "Documentation 📚",
      "type": "PropertyValue",
      "value": "<a href=\"https://docs.gotosocial.org\" rel=\"nofollow noreferrer noopener\" target=\"_blank\">https://docs.gotosocial.org</a>"
    },
    {
      "name": "Give us money if you wanna 😎",
      "type": "PropertyValue",
      "value": "<a href=\"https://liberapay.com/gotosocial\" rel=\"nofollow noreferrer noopener\" target=\"_blank\">https://liberapay.com/gotosocial</a><br><a href=\"https://opencollective.com/gotosocial\" rel=\"nofollow noreferrer noopener\" target=\"_blank\">https://opencollective.com/gotosocial</a>"
    },
    {
      "name": "Matrix space + help channel 💬",
      "type": "PropertyValue",
      "value": "<a href=\"https://matrix.to/#/%23gotosocial-space:superseriousbusiness.org\" rel=\"nofollow noreferrer noopener\" target=\"_blank\">https://matrix.to/#/#gotosocial-space:superseriousbusiness.org</a>"
    },
    {
      "name": "Code :win3_setup:",
      "type": "PropertyValue",
      "value": "<a href=\"https://github.com/superseriousbusiness/gotosocial\" rel=\"nofollow noreferrer noopener\" target=\"_blank\">https://github.com/superseriousbusiness/gotosocial</a><br><a href=\"https://codeberg.org/superseriousbusiness/gotosocial\" rel=\"nofollow noreferrer noopener\" target=\"_blank\">https://codeberg.org/superseriousbusiness/gotosocial</a>"
    }
  ],
  "discoverable": true,
  "featured": "https://gts.superseriousbusiness.org/users/gotosocial/collections/featured",
  "followers": "https://gts.superseriousbusiness.org/users/gotosocial/followers",
  "following": "https://gts.superseriousbusiness.org/users/gotosocial/following",
  "icon": {
    "mediaType": "image/png",
    "type": "Image",
    "url": "https://gts.superseriousbusiness.org/fileserver/016VP4S3BP3QSJBACEQ6J1VFDX/attachment/original/01N7F4KRTAWQP5NTQTHQ56XHD7.png"
  },
  "id": "https://gts.superseriousbusiness.org/users/gotosocial",
  "image": {
    "mediaType": "image/png",
    "type": "Image",
    "url": "https://gts.superseriousbusiness.org/fileserver/016VP4S3BP3QSJBACEQ6J1VFDX/attachment/original/01HVXG7ZB92RBAQ2GYRTK0T0Y9.png"
  },
  "inbox": "https://gts.superseriousbusiness.org/users/gotosocial/inbox",
  "manuallyApprovesFollowers": false,
  "name": "GoToSocial",
  "outbox": "https://gts.superseriousbusiness.org/users/gotosocial/outbox",
  "preferredUsername": "gotosocial",
  "publicKey": {
    "id": "https://gts.superseriousbusiness.org/users/gotosocial/main-key",
    "owner": "https://gts.superseriousbusiness.org/users/gotosocial",
    "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzGB3yDvMl+8p+ViutVRG\nVDl9FO7ZURYXnwB3TedSfG13jyskoiMDNvsbLoUQM9ajZPB0zxJPZUlB/W3BWHRC\nNFQglE5DkB30GjTClNZoOrx64vLRT5wAEwIOjklKVNk9GJi1hFFxrgj931WtxyML\nBvo+TdEblBcoru6MKAov8IU4JjQj5KUmjnW12Rox8dj/rfGtdaH8uJ14vLgvlrAb\neQbN5Ghaxh9DGTo1337O9a9qOsir8YQqazl8ahzS2gvYleV+ou09RDhS75q9hdF2\nLI+1IvFEQ2ZO2tLk3umUP1ioa+5CWKsWD0GAXbQu9uunAV0VoExP4+/9WYOuP0ei\nKwIDAQAB\n-----END PUBLIC KEY-----\n"
  },
  "summary": "<p>The official account for GoToSocial news, updates, release announcements, and so on!</p><p><strong>This is not a support account! Support questions addressed to this account will go unanswered. For support please see the links!</strong></p><p>GoToSocial is an open-source, ActivityPub compatible, federated social media server. It's easy to install, lightweight, and runs quite well on small servers 🦥</p><p>GoToSocial is still <strong>alpha software</strong>. Your support is appreciated as we fix bugs, add features, and move towards beta ❤️</p>",
  "tag": {
    "icon": {
      "mediaType": "image/png",
      "type": "Image",
      "url": "https://gts.superseriousbusiness.org/fileserver/0172YT4H04ZHDY4SPHHNJQEQSM/emoji/original/01QCMZZX058CAPY1AZSNZ6NK7B.png"
    },
    "id": "https://gts.superseriousbusiness.org/emoji/01QCMZZX058CAPY1AZSNZ6NK7B",
    "name": ":win3_setup:",
    "type": "Emoji",
    "updated": "2022-10-09T14:29:12+02:00"
  },
  "type": "Person",
  "url": "https://gts.superseriousbusiness.org/@gotosocial"
}`

var profileJSONOut = `{
  "@context": [
    "https://w3id.org/security/v1",
    "https://www.w3.org/ns/activitystreams",
    {
      "Emoji": "toot:Emoji",
      "PropertyValue": "schema:PropertyValue",
      "discoverable": "toot:discoverable",
      "featured": {
        "@id": "toot:featured",
        "@type": "@id"
      },
      "manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
      "schema": "http://schema.org#",
      "toot": "http://joinmastodon.org/ns#",
      "value": "schema:value"
    }
  ],
  "attachment": [
    {
      "name": "Documentation 📚",
      "type": "PropertyValue",
      "value": "\u003ca href=\"https://docs.gotosocial.org\" rel=\"nofollow noreferrer noopener\" target=\"_blank\"\u003ehttps://docs.gotosocial.org\u003c/a\u003e"
    },
    {
      "name": "Give us money if you wanna 😎",
      "type": "PropertyValue",
      "value": "\u003ca href=\"https://liberapay.com/gotosocial\" rel=\"nofollow noreferrer noopener\" target=\"_blank\"\u003ehttps://liberapay.com/gotosocial\u003c/a\u003e\u003cbr\u003e\u003ca href=\"https://opencollective.com/gotosocial\" rel=\"nofollow noreferrer noopener\" target=\"_blank\"\u003ehttps://opencollective.com/gotosocial\u003c/a\u003e"
    },
    {
      "name": "Matrix space + help channel 💬",
      "type": "PropertyValue",
      "value": "\u003ca href=\"https://matrix.to/#/%23gotosocial-space:superseriousbusiness.org\" rel=\"nofollow noreferrer noopener\" target=\"_blank\"\u003ehttps://matrix.to/#/#gotosocial-space:superseriousbusiness.org\u003c/a\u003e"
    },
    {
      "name": "Code :win3_setup:",
      "type": "PropertyValue",
      "value": "\u003ca href=\"https://github.com/superseriousbusiness/gotosocial\" rel=\"nofollow noreferrer noopener\" target=\"_blank\"\u003ehttps://github.com/superseriousbusiness/gotosocial\u003c/a\u003e\u003cbr\u003e\u003ca href=\"https://codeberg.org/superseriousbusiness/gotosocial\" rel=\"nofollow noreferrer noopener\" target=\"_blank\"\u003ehttps://codeberg.org/superseriousbusiness/gotosocial\u003c/a\u003e"
    }
  ],
  "discoverable": true,
  "featured": "https://gts.superseriousbusiness.org/users/gotosocial/collections/featured",
  "followers": "https://gts.superseriousbusiness.org/users/gotosocial/followers",
  "following": "https://gts.superseriousbusiness.org/users/gotosocial/following",
  "icon": {
    "mediaType": "image/png",
    "type": "Image",
    "url": "https://gts.superseriousbusiness.org/fileserver/016VP4S3BP3QSJBACEQ6J1VFDX/attachment/original/01N7F4KRTAWQP5NTQTHQ56XHD7.png"
  },
  "id": "https://gts.superseriousbusiness.org/users/gotosocial",
  "image": {
    "mediaType": "image/png",
    "type": "Image",
    "url": "https://gts.superseriousbusiness.org/fileserver/016VP4S3BP3QSJBACEQ6J1VFDX/attachment/original/01HVXG7ZB92RBAQ2GYRTK0T0Y9.png"
  },
  "inbox": "https://gts.superseriousbusiness.org/users/gotosocial/inbox",
  "manuallyApprovesFollowers": false,
  "name": "GoToSocial",
  "outbox": "https://gts.superseriousbusiness.org/users/gotosocial/outbox",
  "preferredUsername": "gotosocial",
  "publicKey": {
    "id": "https://gts.superseriousbusiness.org/users/gotosocial/main-key",
    "owner": "https://gts.superseriousbusiness.org/users/gotosocial",
    "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzGB3yDvMl+8p+ViutVRG\nVDl9FO7ZURYXnwB3TedSfG13jyskoiMDNvsbLoUQM9ajZPB0zxJPZUlB/W3BWHRC\nNFQglE5DkB30GjTClNZoOrx64vLRT5wAEwIOjklKVNk9GJi1hFFxrgj931WtxyML\nBvo+TdEblBcoru6MKAov8IU4JjQj5KUmjnW12Rox8dj/rfGtdaH8uJ14vLgvlrAb\neQbN5Ghaxh9DGTo1337O9a9qOsir8YQqazl8ahzS2gvYleV+ou09RDhS75q9hdF2\nLI+1IvFEQ2ZO2tLk3umUP1ioa+5CWKsWD0GAXbQu9uunAV0VoExP4+/9WYOuP0ei\nKwIDAQAB\n-----END PUBLIC KEY-----\n"
  },
  "summary": "\u003cp\u003eThe official account for GoToSocial news, updates, release announcements, and so on!\u003c/p\u003e\u003cp\u003e\u003cstrong\u003eThis is not a support account! Support questions addressed to this account will go unanswered. For support please see the links!\u003c/strong\u003e\u003c/p\u003e\u003cp\u003eGoToSocial is an open-source, ActivityPub compatible, federated social media server. It's easy to install, lightweight, and runs quite well on small servers 🦥\u003c/p\u003e\u003cp\u003eGoToSocial is still \u003cstrong\u003ealpha software\u003c/strong\u003e. Your support is appreciated as we fix bugs, add features, and move towards beta ❤️\u003c/p\u003e",
  "tag": {
    "icon": {
      "mediaType": "image/png",
      "type": "Image",
      "url": "https://gts.superseriousbusiness.org/fileserver/0172YT4H04ZHDY4SPHHNJQEQSM/emoji/original/01QCMZZX058CAPY1AZSNZ6NK7B.png"
    },
    "id": "https://gts.superseriousbusiness.org/emoji/01QCMZZX058CAPY1AZSNZ6NK7B",
    "name": ":win3_setup:",
    "type": "Emoji",
    "updated": "2022-10-09T14:29:12+02:00"
  },
  "type": "Person",
  "url": "https://gts.superseriousbusiness.org/@gotosocial"
}`

var postJSONIn = `{
  "@context": [
    "https://www.w3.org/ns/activitystreams",
    "http://joinmastodon.org/ns"
  ],
  "attachment": [
    {
      "blurhash": "L79@uB.AgQtm%MogwI#+VDs;t7j=",
      "mediaType": "image/png",
      "name": "GoToSocial settings panel showing a tab-panel menu where different interaction types can be permitted or denied per visibility level.",
      "type": "Document",
      "url": "https://gts.superseriousbusiness.org/fileserver/016VP4S3BP3QSJBACEQ6J1VFDX/attachment/original/01J3ZKTT7MM7E5TCQCR1T7FSN2.png"
    }
  ],
  "attributedTo": "https://gts.superseriousbusiness.org/users/gotosocial",
  "cc": "https://gts.superseriousbusiness.org/users/gotosocial/followers",
  "content": "<p>If you follow any of the <a href=\"https://gts.superseriousbusiness.org/tags/gotosocial\" class=\"mention hashtag\" rel=\"tag nofollow noreferrer noopener\" target=\"_blank\">#<span>GoToSocial</span></a> developers you've probably seen this going around already, but 0.17.0 of <a href=\"https://gts.superseriousbusiness.org/tags/gotosocial\" class=\"mention hashtag\" rel=\"tag nofollow noreferrer noopener\" target=\"_blank\">#<span>GoToSocial</span></a> will be the first release that includes interaction policies, aka reply-controls.</p><p>In the first iteration of this feature, you'll be able to configure your account so that new posts created by you will have an interaction policy set on them, which determines whether your instance drops or accepts replies, likes, and boosts of your posts, depending on the visibility of the post, and whether or not an account trying to interact with you is in your followers/following list.</p><p>So for example, you will be able to create Public posts that can only be replied to by your followers and people you follow, or unlisted posts that nobody can reply to or like, etc.</p><p>GoToSocial interaction policies will be a superset of other reply control proposals created elsewhere (and already implemented by softwares like Pixelfed and Peertube), so your GoToSocial instance should recognize interaction restrictions set not only by other GoToSocial instances, but by Pixelfed and Peertube as well.</p><p>If you're interested in reading about how this will work on a protocol level, you can take a look at the documentation here: <a href=\"https://docs.gotosocial.org/en/latest/federation/posts/#interaction-policy\" rel=\"nofollow noreferrer noopener\" target=\"_blank\">https://docs.gotosocial.org/en/latest/federation/posts/#interaction-policy</a></p><p>Please note that this feature is not 100% finished yet, and may be subject to change before release. We're aware of where the headaches and difficulties are, so please don't reply to this post griping about them; we already know (and this instance is still running on 0.16.0 so no interaction policies yet).</p><p>Thanks for reading :)</p>",
  "contentMap": {
    "EN": "<p>If you follow any of the <a href=\"https://gts.superseriousbusiness.org/tags/gotosocial\" class=\"mention hashtag\" rel=\"tag nofollow noreferrer noopener\" target=\"_blank\">#<span>GoToSocial</span></a> developers you've probably seen this going around already, but 0.17.0 of <a href=\"https://gts.superseriousbusiness.org/tags/gotosocial\" class=\"mention hashtag\" rel=\"tag nofollow noreferrer noopener\" target=\"_blank\">#<span>GoToSocial</span></a> will be the first release that includes interaction policies, aka reply-controls.</p><p>In the first iteration of this feature, you'll be able to configure your account so that new posts created by you will have an interaction policy set on them, which determines whether your instance drops or accepts replies, likes, and boosts of your posts, depending on the visibility of the post, and whether or not an account trying to interact with you is in your followers/following list.</p><p>So for example, you will be able to create Public posts that can only be replied to by your followers and people you follow, or unlisted posts that nobody can reply to or like, etc.</p><p>GoToSocial interaction policies will be a superset of other reply control proposals created elsewhere (and already implemented by softwares like Pixelfed and Peertube), so your GoToSocial instance should recognize interaction restrictions set not only by other GoToSocial instances, but by Pixelfed and Peertube as well.</p><p>If you're interested in reading about how this will work on a protocol level, you can take a look at the documentation here: <a href=\"https://docs.gotosocial.org/en/latest/federation/posts/#interaction-policy\" rel=\"nofollow noreferrer noopener\" target=\"_blank\">https://docs.gotosocial.org/en/latest/federation/posts/#interaction-policy</a></p><p>Please note that this feature is not 100% finished yet, and may be subject to change before release. We're aware of where the headaches and difficulties are, so please don't reply to this post griping about them; we already know (and this instance is still running on 0.16.0 so no interaction policies yet).</p><p>Thanks for reading :)</p>"
  },
  "id": "https://gts.superseriousbusiness.org/users/gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC",
  "published": "2024-07-29T18:12:01+02:00",
  "replies": {
    "first": {
      "id": "https://gts.superseriousbusiness.org/users/gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC/replies?page=true",
      "next": "https://gts.superseriousbusiness.org/users/gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC/replies?only_other_accounts=false&page=true",
      "partOf": "https://gts.superseriousbusiness.org/users/gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC/replies",
      "type": "CollectionPage"
    },
    "id": "https://gts.superseriousbusiness.org/users/gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC/replies",
    "type": "Collection"
  },
  "sensitive": false,
  "summary": "",
  "tag": {
    "href": "https://gts.superseriousbusiness.org/tags/gotosocial",
    "name": "#gotosocial",
    "type": "Hashtag"
  },
  "to": "https://www.w3.org/ns/activitystreams#Public",
  "type": "Note",
  "url": "https://gts.superseriousbusiness.org/@gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC"
}`

var postJSONOut = `{
  "@context": [
    "https://www.w3.org/ns/activitystreams",
    {
      "Hashtag": "as:Hashtag",
      "blurhash": "toot:blurhash",
      "sensitive": "as:sensitive",
      "toot": "http://joinmastodon.org/ns#"
    }
  ],
  "attachment": {
    "blurhash": "L79@uB.AgQtm%MogwI#+VDs;t7j=",
    "mediaType": "image/png",
    "name": "GoToSocial settings panel showing a tab-panel menu where different interaction types can be permitted or denied per visibility level.",
    "type": "Document",
    "url": "https://gts.superseriousbusiness.org/fileserver/016VP4S3BP3QSJBACEQ6J1VFDX/attachment/original/01J3ZKTT7MM7E5TCQCR1T7FSN2.png"
  },
  "attributedTo": "https://gts.superseriousbusiness.org/users/gotosocial",
  "cc": "https://gts.superseriousbusiness.org/users/gotosocial/followers",
  "content": "\u003cp\u003eIf you follow any of the \u003ca href=\"https://gts.superseriousbusiness.org/tags/gotosocial\" class=\"mention hashtag\" rel=\"tag nofollow noreferrer noopener\" target=\"_blank\"\u003e#\u003cspan\u003eGoToSocial\u003c/span\u003e\u003c/a\u003e developers you've probably seen this going around already, but 0.17.0 of \u003ca href=\"https://gts.superseriousbusiness.org/tags/gotosocial\" class=\"mention hashtag\" rel=\"tag nofollow noreferrer noopener\" target=\"_blank\"\u003e#\u003cspan\u003eGoToSocial\u003c/span\u003e\u003c/a\u003e will be the first release that includes interaction policies, aka reply-controls.\u003c/p\u003e\u003cp\u003eIn the first iteration of this feature, you'll be able to configure your account so that new posts created by you will have an interaction policy set on them, which determines whether your instance drops or accepts replies, likes, and boosts of your posts, depending on the visibility of the post, and whether or not an account trying to interact with you is in your followers/following list.\u003c/p\u003e\u003cp\u003eSo for example, you will be able to create Public posts that can only be replied to by your followers and people you follow, or unlisted posts that nobody can reply to or like, etc.\u003c/p\u003e\u003cp\u003eGoToSocial interaction policies will be a superset of other reply control proposals created elsewhere (and already implemented by softwares like Pixelfed and Peertube), so your GoToSocial instance should recognize interaction restrictions set not only by other GoToSocial instances, but by Pixelfed and Peertube as well.\u003c/p\u003e\u003cp\u003eIf you're interested in reading about how this will work on a protocol level, you can take a look at the documentation here: \u003ca href=\"https://docs.gotosocial.org/en/latest/federation/posts/#interaction-policy\" rel=\"nofollow noreferrer noopener\" target=\"_blank\"\u003ehttps://docs.gotosocial.org/en/latest/federation/posts/#interaction-policy\u003c/a\u003e\u003c/p\u003e\u003cp\u003ePlease note that this feature is not 100% finished yet, and may be subject to change before release. We're aware of where the headaches and difficulties are, so please don't reply to this post griping about them; we already know (and this instance is still running on 0.16.0 so no interaction policies yet).\u003c/p\u003e\u003cp\u003eThanks for reading :)\u003c/p\u003e",
  "id": "https://gts.superseriousbusiness.org/users/gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC",
  "published": "2024-07-29T18:12:01+02:00",
  "replies": {
    "first": {
      "id": "https://gts.superseriousbusiness.org/users/gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC/replies?page=true",
      "next": "https://gts.superseriousbusiness.org/users/gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC/replies?only_other_accounts=false\u0026page=true",
      "partOf": "https://gts.superseriousbusiness.org/users/gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC/replies",
      "type": "CollectionPage"
    },
    "id": "https://gts.superseriousbusiness.org/users/gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC/replies",
    "type": "Collection"
  },
  "sensitive": false,
  "summary": "",
  "tag": {
    "href": "https://gts.superseriousbusiness.org/tags/gotosocial",
    "name": "#gotosocial",
    "type": "Hashtag"
  },
  "to": "https://www.w3.org/ns/activitystreams#Public",
  "type": "Note",
  "url": "https://gts.superseriousbusiness.org/@gotosocial/statuses/01J3ZM6N4VQ1NS60RHGAVVWMFC"
}`

var outboxJSONIn = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "first": "https://gts.superseriousbusiness.org/users/gotosocial/outbox?limit=40",
  "id": "https://gts.superseriousbusiness.org/users/gotosocial/outbox",
  "totalItems": 951,
  "type": "OrderedCollection"
}`

var outboxJSONOut = `{
  "@context": "https://www.w3.org/ns/activitystreams",
  "first": "https://gts.superseriousbusiness.org/users/gotosocial/outbox?limit=40",
  "id": "https://gts.superseriousbusiness.org/users/gotosocial/outbox",
  "totalItems": 951,
  "type": "OrderedCollection"
}`

func TestSerializeContext(t *testing.T) {
	type testCase struct {
		in            vocab.Type
		expectOutJSON string
	}

	for _, tc := range []testCase{
		{
			in:            JSONStringToType(profileJSONIn),
			expectOutJSON: profileJSONOut,
		},
		{
			in:            JSONStringToType(postJSONIn),
			expectOutJSON: postJSONOut,
		},
		{
			in:            JSONStringToType(outboxJSONIn),
			expectOutJSON: outboxJSONOut,
		},
	} {
		out, err := streams.Serialize(tc.in)
		if err != nil {
			t.Logf("expected no error, got %v", err)
			t.Fail()
			continue
		}

		outJSON, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			t.Logf("expected no error, got %v", err)
			t.Fail()
			continue
		}

		if string(outJSON) != tc.expectOutJSON {
			t.Logf("expected json:\n%s", tc.expectOutJSON)
			t.Logf("got json:\n%s", outJSON)
			t.Fail()
			continue
		}
	}
}

func JSONStringToType(in string) vocab.Type {
	m := make(map[string]any)
	if err := json.Unmarshal([]byte(in), &m); err != nil {
		panic(err)
	}

	t, err := streams.ToType(context.Background(), m)
	if err != nil {
		panic(err)
	}

	return t
}
