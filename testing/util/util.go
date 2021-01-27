package util

import (
	"github.com/jarcoal/httpmock"
	"github.com/muhammednagy/pipedrive-challenge/db"
	"github.com/muhammednagy/pipedrive-challenge/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

func TearDown(dbConnection *gorm.DB) {
	if err := db.DropTestDB(dbConnection); err != nil {
		log.Fatal(err)
	}
}

func MockPipedrive() {
	httpmock.Activate()
	// Mock creating activity
	httpmock.RegisterResponder("POST", "https://api.pipedrive.com/v1/activities",
		httpmock.NewStringResponder(200,
			`{ "success": true, "data": { "id": 3, "company_id": 7815951, "user_id": 11947605, "done": false, "type": "call", "reference_type": null, "reference_id": null, "conference_meeting_client": null, "conference_meeting_url": null, "due_date": "2021-01-21", "due_time": "", "duration": "", "busy_flag": null, "add_time": "2021-01-21 20:26:29", "marked_as_done_time": "", "last_notification_time": null, "last_notification_user_id": null, "notification_language_id": null, "subject": "test activity", "public_description": null, "calendar_sync_include_context": null, "location": null, "org_id": null, "person_id": 3, "deal_id": null, "lead_id": null, "lead_title": "", "active_flag": true, "update_time": "2021-01-21 20:26:29", "update_user_id": null, "gcal_event_id": null, "google_calendar_id": null, "google_calendar_etag": null, "source_timezone": null, "rec_rule": null, "rec_rule_extension": null, "rec_master_activity_id": null, "conference_meeting_id": null, "note": "test note", "created_by_user_id": 11947605, "location_subpremise": null, "location_street_number": null, "location_route": null, "location_sublocality": null, "location_locality": null, "location_admin_area_level_1": null, "location_admin_area_level_2": null, "location_country": null, "location_postal_code": null, "location_formatted_address": null, "attendees": null, "participants": [ { "person_id": 3, "primary_flag": true } ], "series": null, "org_name": null, "person_name": "test", "deal_title": null, "owner_name": "Nagy", "person_dropbox_bcc": "pipedrivetest-sandbox2@pipedrivemail.com", "deal_dropbox_bcc": null, "assigned_to_user_id": 11947605, "type_name": "Call", "file": null }, "additional_data": { "updates_story_id": 8 }, "related_objects": { "person": { "3": { "active_flag": true, "id": 3, "name": "test", "email": [ { "value": "", "primary": true } ], "phone": [ { "value": "", "primary": true } ] } }, "user": { "11947605": { "id": 11947605, "name": "Nagy", "email": "me@muhnagy.com", "has_pic": 0, "pic_hash": null, "active_flag": true } } } }`,
		))
	// Mock creating activity
	httpmock.RegisterResponder("POST", "https://api.pipedrive.com/v1/persons",
		httpmock.NewStringResponder(200,
			`{ "success": true, "data": { "id": 9, "company_id": 7815951, "owner_id": { "id": 11947605, "name": "Nagy", "email": "me@muhnagy.com", "has_pic": 0, "pic_hash": null, "active_flag": true, "value": 11947605 }, "org_id": null, "name": "muhammednagy3", "first_name": "muhammednagy3", "last_name": null, "open_deals_count": 0, "related_open_deals_count": 0, "closed_deals_count": 0, "related_closed_deals_count": 0, "participant_open_deals_count": 0, "participant_closed_deals_count": 0, "email_messages_count": 0, "activities_count": 0, "done_activities_count": 0, "undone_activities_count": 0, "files_count": 0, "notes_count": 0, "followers_count": 0, "won_deals_count": 0, "related_won_deals_count": 0, "lost_deals_count": 0, "related_lost_deals_count": 0, "active_flag": true, "phone": [ { "value": "", "primary": true } ], "email": [ { "value": "", "primary": true } ], "first_char": "m", "update_time": "2021-01-23 20:27:36", "add_time": "2021-01-23 20:27:36", "visible_to": "3", "picture_id": null, "next_activity_date": null, "next_activity_time": null, "next_activity_id": null, "last_activity_id": null, "last_activity_date": null, "last_incoming_mail_time": null, "last_outgoing_mail_time": null, "label": null, "org_name": null, "cc_email": "pipedrivetest-sandbox2@pipedrivemail.com", "owner_name": "Nagy" }, "related_objects": { "user": { "11947605": { "id": 11947605, "name": "Nagy", "email": "me@muhnagy.com", "has_pic": 0, "pic_hash": null, "active_flag": true } } } }`))
}

func MockGithub() {
	httpmock.Activate()
	// Mock getting gists
	httpmock.RegisterResponder("GET", "https://api.github.com/users/muhammednagy/gists",
		httpmock.NewStringResponder(200,
			`[{"url":"https://api.github.com/gists/9ab5f20c69335e4917575333e3b9e8e0","forks_url":"https://api.github.com/gists/9ab5f20c69335e4917575333e3b9e8e0/forks","commits_url":"https://api.github.com/gists/9ab5f20c69335e4917575333e3b9e8e0/commits","id":"9ab5f20c69335e4917575333e3b9e8e0","node_id":"MDQ6R2lzdDlhYjVmMjBjNjkzMzVlNDkxNzU3NTMzM2UzYjllOGUw","git_pull_url":"https://gist.github.com/9ab5f20c69335e4917575333e3b9e8e0.git","git_push_url":"https://gist.github.com/9ab5f20c69335e4917575333e3b9e8e0.git","html_url":"https://gist.github.com/9ab5f20c69335e4917575333e3b9e8e0","files":{"test2.rb":{"filename":"test2.rb","type":"application/x-ruby","language":"Ruby","raw_url":"https://gist.githubusercontent.com/muhammednagy/9ab5f20c69335e4917575333e3b9e8e0/raw/a83efbf55c1f69ea22ba6c7cb6cfe58f5a84ef35/test2.rb","size":17}},"public":true,"created_at":"2021-01-23T03:03:38Z","updated_at":"2021-01-23T03:03:48Z","description":"","comments":0,"user":null,"comments_url":"https://api.github.com/gists/9ab5f20c69335e4917575333e3b9e8e0/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false},{"url":"https://api.github.com/gists/d4ac65be4da7f3fdd890f6275030107b","forks_url":"https://api.github.com/gists/d4ac65be4da7f3fdd890f6275030107b/forks","commits_url":"https://api.github.com/gists/d4ac65be4da7f3fdd890f6275030107b/commits","id":"d4ac65be4da7f3fdd890f6275030107b","node_id":"MDQ6R2lzdGQ0YWM2NWJlNGRhN2YzZmRkODkwZjYyNzUwMzAxMDdi","git_pull_url":"https://gist.github.com/d4ac65be4da7f3fdd890f6275030107b.git","git_push_url":"https://gist.github.com/d4ac65be4da7f3fdd890f6275030107b.git","html_url":"https://gist.github.com/d4ac65be4da7f3fdd890f6275030107b","files":{"test.rb":{"filename":"test.rb","type":"application/x-ruby","language":"Ruby","raw_url":"https://gist.githubusercontent.com/muhammednagy/d4ac65be4da7f3fdd890f6275030107b/raw/46013d8150261760982bd73cb3d1f3e8676fb2df/test.rb","size":16}},"public":true,"created_at":"2021-01-23T02:36:10Z","updated_at":"2021-01-23T02:36:10Z","description":"test","comments":0,"user":null,"comments_url":"https://api.github.com/gists/d4ac65be4da7f3fdd890f6275030107b/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false},{"url":"https://api.github.com/gists/b31dea4f68e10236176506785ad29cfb","forks_url":"https://api.github.com/gists/b31dea4f68e10236176506785ad29cfb/forks","commits_url":"https://api.github.com/gists/b31dea4f68e10236176506785ad29cfb/commits","id":"b31dea4f68e10236176506785ad29cfb","node_id":"MDQ6R2lzdGIzMWRlYTRmNjhlMTAyMzYxNzY1MDY3ODVhZDI5Y2Zi","git_pull_url":"https://gist.github.com/b31dea4f68e10236176506785ad29cfb.git","git_push_url":"https://gist.github.com/b31dea4f68e10236176506785ad29cfb.git","html_url":"https://gist.github.com/b31dea4f68e10236176506785ad29cfb","files":{"gistfile1.txt":{"filename":"gistfile1.txt","type":"text/plain","language":"Text","raw_url":"https://gist.githubusercontent.com/muhammednagy/b31dea4f68e10236176506785ad29cfb/raw/6f35899d341015e83bba3c0a7e53aae471fffdf8/gistfile1.txt","size":42}},"public":true,"created_at":"2017-10-04T20:19:51Z","updated_at":"2017-10-04T20:19:51Z","description":"","comments":0,"user":null,"comments_url":"https://api.github.com/gists/b31dea4f68e10236176506785ad29cfb/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false},{"url":"https://api.github.com/gists/c5cfd9b600125afba0f3cdf1ff1c3616","forks_url":"https://api.github.com/gists/c5cfd9b600125afba0f3cdf1ff1c3616/forks","commits_url":"https://api.github.com/gists/c5cfd9b600125afba0f3cdf1ff1c3616/commits","id":"c5cfd9b600125afba0f3cdf1ff1c3616","node_id":"MDQ6R2lzdGM1Y2ZkOWI2MDAxMjVhZmJhMGYzY2RmMWZmMWMzNjE2","git_pull_url":"https://gist.github.com/c5cfd9b600125afba0f3cdf1ff1c3616.git","git_push_url":"https://gist.github.com/c5cfd9b600125afba0f3cdf1ff1c3616.git","html_url":"https://gist.github.com/c5cfd9b600125afba0f3cdf1ff1c3616","files":{"gistfile1.txt":{"filename":"gistfile1.txt","type":"text/plain","language":"Text","raw_url":"https://gist.githubusercontent.com/muhammednagy/c5cfd9b600125afba0f3cdf1ff1c3616/raw/ac13a8b89e149a5cf8de88c0921f7067c2fbf2b6/gistfile1.txt","size":42}},"public":true,"created_at":"2017-10-04T19:50:22Z","updated_at":"2017-10-04T19:50:22Z","description":"","comments":0,"user":null,"comments_url":"https://api.github.com/gists/c5cfd9b600125afba0f3cdf1ff1c3616/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false},{"url":"https://api.github.com/gists/d8147d7a516b24d63b8ddfef7ab317fc","forks_url":"https://api.github.com/gists/d8147d7a516b24d63b8ddfef7ab317fc/forks","commits_url":"https://api.github.com/gists/d8147d7a516b24d63b8ddfef7ab317fc/commits","id":"d8147d7a516b24d63b8ddfef7ab317fc","node_id":"MDQ6R2lzdGQ4MTQ3ZDdhNTE2YjI0ZDYzYjhkZGZlZjdhYjMxN2Zj","git_pull_url":"https://gist.github.com/d8147d7a516b24d63b8ddfef7ab317fc.git","git_push_url":"https://gist.github.com/d8147d7a516b24d63b8ddfef7ab317fc.git","html_url":"https://gist.github.com/d8147d7a516b24d63b8ddfef7ab317fc","files":{"my eth test address":{"filename":"my eth test address","type":"text/plain","language":null,"raw_url":"https://gist.githubusercontent.com/muhammednagy/d8147d7a516b24d63b8ddfef7ab317fc/raw/ac13a8b89e149a5cf8de88c0921f7067c2fbf2b6/my%20eth%20test%20address","size":42}},"public":true,"created_at":"2017-10-04T01:43:24Z","updated_at":"2017-10-04T01:43:25Z","description":"","comments":0,"user":null,"comments_url":"https://api.github.com/gists/d8147d7a516b24d63b8ddfef7ab317fc/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false},{"url":"https://api.github.com/gists/e2ac3e576216ebf576664f43e2f60002","forks_url":"https://api.github.com/gists/e2ac3e576216ebf576664f43e2f60002/forks","commits_url":"https://api.github.com/gists/e2ac3e576216ebf576664f43e2f60002/commits","id":"e2ac3e576216ebf576664f43e2f60002","node_id":"MDQ6R2lzdGUyYWMzZTU3NjIxNmViZjU3NjY2NGY0M2UyZjYwMDAy","git_pull_url":"https://gist.github.com/e2ac3e576216ebf576664f43e2f60002.git","git_push_url":"https://gist.github.com/e2ac3e576216ebf576664f43e2f60002.git","html_url":"https://gist.github.com/e2ac3e576216ebf576664f43e2f60002","files":{"add-new-crypto-to-peatio.md":{"filename":"add-new-crypto-to-peatio.md","type":"text/markdown","language":"Markdown","raw_url":"https://gist.githubusercontent.com/muhammednagy/e2ac3e576216ebf576664f43e2f60002/raw/d5d350a229fe44c98c9d086e5b57fe20016f6238/add-new-crypto-to-peatio.md","size":1956}},"public":true,"created_at":"2017-07-22T22:01:22Z","updated_at":"2019-01-07T16:38:15Z","description":"Adding A New Cryptocurrency to Peatio","comments":4,"user":null,"comments_url":"https://api.github.com/gists/e2ac3e576216ebf576664f43e2f60002/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false},{"url":"https://api.github.com/gists/47f357f7c429cf9a6fe80431ee778e70","forks_url":"https://api.github.com/gists/47f357f7c429cf9a6fe80431ee778e70/forks","commits_url":"https://api.github.com/gists/47f357f7c429cf9a6fe80431ee778e70/commits","id":"47f357f7c429cf9a6fe80431ee778e70","node_id":"MDQ6R2lzdDQ3ZjM1N2Y3YzQyOWNmOWE2ZmU4MDQzMWVlNzc4ZTcw","git_pull_url":"https://gist.github.com/47f357f7c429cf9a6fe80431ee778e70.git","git_push_url":"https://gist.github.com/47f357f7c429cf9a6fe80431ee778e70.git","html_url":"https://gist.github.com/47f357f7c429cf9a6fe80431ee778e70","files":{"fasterdownloadwithwget.sh":{"filename":"fasterdownloadwithwget.sh","type":"application/x-sh","language":"Shell","raw_url":"https://gist.githubusercontent.com/muhammednagy/47f357f7c429cf9a6fe80431ee778e70/raw/27984bcc107bd51e38e044760a91b5feb4bcc01a/fasterdownloadwithwget.sh","size":188}},"public":true,"created_at":"2017-03-31T19:37:54Z","updated_at":"2017-03-31T19:37:54Z","description":"download bunch of files quietly from a file ","comments":0,"user":null,"comments_url":"https://api.github.com/gists/47f357f7c429cf9a6fe80431ee778e70/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false},{"url":"https://api.github.com/gists/d525fe701c9c064d188bffc25b611029","forks_url":"https://api.github.com/gists/d525fe701c9c064d188bffc25b611029/forks","commits_url":"https://api.github.com/gists/d525fe701c9c064d188bffc25b611029/commits","id":"d525fe701c9c064d188bffc25b611029","node_id":"MDQ6R2lzdGQ1MjVmZTcwMWM5YzA2NGQxODhiZmZjMjViNjExMDI5","git_pull_url":"https://gist.github.com/d525fe701c9c064d188bffc25b611029.git","git_push_url":"https://gist.github.com/d525fe701c9c064d188bffc25b611029.git","html_url":"https://gist.github.com/d525fe701c9c064d188bffc25b611029","files":{"assignment.erl":{"filename":"assignment.erl","type":"text/plain","language":"Erlang","raw_url":"https://gist.githubusercontent.com/muhammednagy/d525fe701c9c064d188bffc25b611029/raw/d16a228feb242fea3362569d733508b35ffd0506/assignment.erl","size":755}},"public":true,"created_at":"2017-02-25T19:12:01Z","updated_at":"2017-02-25T19:12:01Z","description":"Week 1 assignment kent university functional programming in erlang","comments":0,"user":null,"comments_url":"https://api.github.com/gists/d525fe701c9c064d188bffc25b611029/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false},{"url":"https://api.github.com/gists/0f1ad3c7999c1bb4052c065e7fb40121","forks_url":"https://api.github.com/gists/0f1ad3c7999c1bb4052c065e7fb40121/forks","commits_url":"https://api.github.com/gists/0f1ad3c7999c1bb4052c065e7fb40121/commits","id":"0f1ad3c7999c1bb4052c065e7fb40121","node_id":"MDQ6R2lzdDBmMWFkM2M3OTk5YzFiYjQwNTJjMDY1ZTdmYjQwMTIx","git_pull_url":"https://gist.github.com/0f1ad3c7999c1bb4052c065e7fb40121.git","git_push_url":"https://gist.github.com/0f1ad3c7999c1bb4052c065e7fb40121.git","html_url":"https://gist.github.com/0f1ad3c7999c1bb4052c065e7fb40121","files":{"bitcount.erl":{"filename":"bitcount.erl","type":"text/plain","language":"Erlang","raw_url":"https://gist.githubusercontent.com/muhammednagy/0f1ad3c7999c1bb4052c065e7fb40121/raw/492a5c5e835def12176a67dfd7ac3e7fc80227f4/bitcount.erl","size":163}},"public":true,"created_at":"2017-02-24T23:27:23Z","updated_at":"2020-12-06T12:16:39Z","description":"count bits in erlang","comments":1,"user":null,"comments_url":"https://api.github.com/gists/0f1ad3c7999c1bb4052c065e7fb40121/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false},{"url":"https://api.github.com/gists/ea75ed93795b5112b82ab5439ccbbe69","forks_url":"https://api.github.com/gists/ea75ed93795b5112b82ab5439ccbbe69/forks","commits_url":"https://api.github.com/gists/ea75ed93795b5112b82ab5439ccbbe69/commits","id":"ea75ed93795b5112b82ab5439ccbbe69","node_id":"MDQ6R2lzdGVhNzVlZDkzNzk1YjUxMTJiODJhYjU0MzljY2JiZTY5","git_pull_url":"https://gist.github.com/ea75ed93795b5112b82ab5439ccbbe69.git","git_push_url":"https://gist.github.com/ea75ed93795b5112b82ab5439ccbbe69.git","html_url":"https://gist.github.com/ea75ed93795b5112b82ab5439ccbbe69","files":{"recursion.erl":{"filename":"recursion.erl","type":"text/plain","language":"Erlang","raw_url":"https://gist.githubusercontent.com/muhammednagy/ea75ed93795b5112b82ab5439ccbbe69/raw/16a1fc30647cbecf0ab3ca823cc641cb0d44ad93/recursion.erl","size":299}},"public":true,"created_at":"2017-02-24T19:38:05Z","updated_at":"2017-02-24T19:38:05Z","description":"Recursion in erlang (Fibonacci, factorial, n dimensions)","comments":0,"user":null,"comments_url":"https://api.github.com/gists/ea75ed93795b5112b82ab5439ccbbe69/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false},{"url":"https://api.github.com/gists/578734582b02cc222104a0213cd7772a","forks_url":"https://api.github.com/gists/578734582b02cc222104a0213cd7772a/forks","commits_url":"https://api.github.com/gists/578734582b02cc222104a0213cd7772a/commits","id":"578734582b02cc222104a0213cd7772a","node_id":"MDQ6R2lzdDU3ODczNDU4MmIwMmNjMjIyMTA0YTAyMTNjZDc3NzJh","git_pull_url":"https://gist.github.com/578734582b02cc222104a0213cd7772a.git","git_push_url":"https://gist.github.com/578734582b02cc222104a0213cd7772a.git","html_url":"https://gist.github.com/578734582b02cc222104a0213cd7772a","files":{"fun.erl":{"filename":"fun.erl","type":"text/plain","language":"Erlang","raw_url":"https://gist.githubusercontent.com/muhammednagy/578734582b02cc222104a0213cd7772a/raw/bd4e7acae6ca6a4ac54a4bd529234b0b237944ce/fun.erl","size":259}},"public":true,"created_at":"2017-02-24T17:13:36Z","updated_at":"2017-02-24T17:54:09Z","description":"variables and patterns in erlang","comments":1,"user":null,"comments_url":"https://api.github.com/gists/578734582b02cc222104a0213cd7772a/comments","owner":{"login":"muhammednagy","id":18580720,"node_id":"MDQ6VXNlcjE4NTgwNzIw","avatar_url":"https://avatars.githubusercontent.com/u/18580720?v=4","gravatar_id":"","url":"https://api.github.com/users/muhammednagy","html_url":"https://github.com/muhammednagy","followers_url":"https://api.github.com/users/muhammednagy/followers","following_url":"https://api.github.com/users/muhammednagy/following{/other_user}","gists_url":"https://api.github.com/users/muhammednagy/gists{/gist_id}","starred_url":"https://api.github.com/users/muhammednagy/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/muhammednagy/subscriptions","organizations_url":"https://api.github.com/users/muhammednagy/orgs","repos_url":"https://api.github.com/users/muhammednagy/repos","events_url":"https://api.github.com/users/muhammednagy/events{/privacy}","received_events_url":"https://api.github.com/users/muhammednagy/received_events","type":"User","site_admin":false},"truncated":false}]`,
		))
}

func LoadFixtures(d *gorm.DB) error {
	p1 := model.Person{
		GithubUsername: "muhammednagy",
	}
	currentTime := time.Now().UTC()
	oneHourBeforeCurrentTime := currentTime.Add(time.Hour * time.Duration(-1))
	oneHourAfterCurrentTime := currentTime.Add(time.Hour * time.Duration(1))
	p2 := model.Person{
		GithubUsername: "muhammednagy2",
		LastVisit:      &currentTime,
		Gists: []model.Gist{{
			DBModel:     model.DBModel{CreatedAt: currentTime},
			Description: "test gist",
			PullURL:     "http://test.dummy/git1",
			Files: []model.GistFile{{
				Name:   "test.go",
				RawURL: "http://test.dummy/test.go",
			}},
		},
			{
				DBModel:     model.DBModel{CreatedAt: oneHourBeforeCurrentTime},
				Description: "test gist made one hour earlier",
				PullURL:     "http://test.dummy/git1",
				Files: []model.GistFile{{
					Name:   "test.go",
					RawURL: "http://test.dummy/test.go",
				}},
			},
			{
				DBModel:     model.DBModel{CreatedAt: oneHourAfterCurrentTime},
				Description: "test gist made one hour later",
				PullURL:     "http://test.dummy/git1",
				Files: []model.GistFile{{
					Name:   "test.go",
					RawURL: "http://test.dummy/test.go",
				}},
			},
		},
	}

	if err := d.Create(&p1).Error; err != nil {
		return err
	}

	return d.Create(&p2).Error
}
