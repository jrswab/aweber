package aweber

type List struct {
	CampaignsCollectionLink             string `json:"campaigns_collection_link"`
	CustomFieldsCollectionLink          string `json:"custom_fields_collection_link"`
	DraftBroadcastsLink                 string `json:"draft_broadcasts_link"`
	ScheduledBroadcastsLink             string `json:"scheduled_broadcasts_link"`
	SentBroadcastsLink                  string `json:"sent_broadcasts_link"`
	HTTPEtag                            string `json:"http_etag"`
	ID                                  int    `json:"id"`
	LandingPagesCollectionLink          string `json:"landing_pages_collection_link"`
	Name                                string `json:"name"`
	ResourceTypeLink                    string `json:"resource_type_link"`
	SegmentsCollectionLink              string `json:"segments_collection_link"`
	SelfLink                            string `json:"self_link"`
	SubscribersCollectionLink           string `json:"subscribers_collection_link"`
	TotalSubscribedSubscribers          int    `json:"total_subscribed_subscribers"`
	TotalSubscribers                    int    `json:"total_subscribers"`
	TotalSubscribersSubscribedToday     int    `json:"total_subscribers_subscribed_today"`
	TotalSubscribersSubscribedYesterday int    `json:"total_subscribers_subscribed_yesterday"`
	TotalUnconfirmedSubscribers         int    `json:"total_unconfirmed_subscribers"`
	TotalUnsubscribedSubscribers        int    `json:"total_unsubscribed_subscribers"`
	UniqueListID                        string `json:"unique_list_id"`
	UUID                                string `json:"uuid"`
	WebFormSplitTestsCollectionLink     string `json:"web_form_split_tests_collection_link"`
	WebFormsCollectionLink              string `json:"web_forms_collection_link"`
}

type Lists struct {
	Entries []struct {
		CampaignsCollectionLink             string `json:"campaigns_collection_link"`
		CustomFieldsCollectionLink          string `json:"custom_fields_collection_link"`
		DraftBroadcastsLink                 string `json:"draft_broadcasts_link"`
		ScheduledBroadcastsLink             string `json:"scheduled_broadcasts_link"`
		SentBroadcastsLink                  string `json:"sent_broadcasts_link"`
		HTTPEtag                            string `json:"http_etag"`
		ID                                  int    `json:"id"`
		LandingPagesCollectionLink          string `json:"landing_pages_collection_link"`
		Name                                string `json:"name"`
		ResourceTypeLink                    string `json:"resource_type_link"`
		SegmentsCollectionLink              string `json:"segments_collection_link"`
		SelfLink                            string `json:"self_link"`
		SubscribersCollectionLink           string `json:"subscribers_collection_link"`
		TotalSubscribedSubscribers          int    `json:"total_subscribed_subscribers"`
		TotalSubscribers                    int    `json:"total_subscribers"`
		TotalSubscribersSubscribedToday     int    `json:"total_subscribers_subscribed_today"`
		TotalSubscribersSubscribedYesterday int    `json:"total_subscribers_subscribed_yesterday"`
		TotalUnconfirmedSubscribers         int    `json:"total_unconfirmed_subscribers"`
		TotalUnsubscribedSubscribers        int    `json:"total_unsubscribed_subscribers"`
		UniqueListID                        string `json:"unique_list_id"`
		UUID                                string `json:"uuid"`
		WebFormSplitTestsCollectionLink     string `json:"web_form_split_tests_collection_link"`
		WebFormsCollectionLink              string `json:"web_forms_collection_link"`
	} `json:"entries"`
	NextCollectionLink string `json:"next_collection_link"`
	PrevCollectionLink string `json:"prev_collection_link"`
	ResourceTypeLink   string `json:"resource_type_link"`
	Start              int    `json:"start"`
	TotalSize          int    `json:"total_size"`
}
