package streams

import (
	propertyaccuracy "github.com/go-fed/activity/streams/impl/activitystreams/property_accuracy"
	propertyactor "github.com/go-fed/activity/streams/impl/activitystreams/property_actor"
	propertyaltitude "github.com/go-fed/activity/streams/impl/activitystreams/property_altitude"
	propertyanyof "github.com/go-fed/activity/streams/impl/activitystreams/property_anyof"
	propertyattachment "github.com/go-fed/activity/streams/impl/activitystreams/property_attachment"
	propertyattributedto "github.com/go-fed/activity/streams/impl/activitystreams/property_attributedto"
	propertyaudience "github.com/go-fed/activity/streams/impl/activitystreams/property_audience"
	propertybcc "github.com/go-fed/activity/streams/impl/activitystreams/property_bcc"
	propertybto "github.com/go-fed/activity/streams/impl/activitystreams/property_bto"
	propertycc "github.com/go-fed/activity/streams/impl/activitystreams/property_cc"
	propertyclosed "github.com/go-fed/activity/streams/impl/activitystreams/property_closed"
	propertycontent "github.com/go-fed/activity/streams/impl/activitystreams/property_content"
	propertycontext "github.com/go-fed/activity/streams/impl/activitystreams/property_context"
	propertycurrent "github.com/go-fed/activity/streams/impl/activitystreams/property_current"
	propertydeleted "github.com/go-fed/activity/streams/impl/activitystreams/property_deleted"
	propertydescribes "github.com/go-fed/activity/streams/impl/activitystreams/property_describes"
	propertyduration "github.com/go-fed/activity/streams/impl/activitystreams/property_duration"
	propertyendtime "github.com/go-fed/activity/streams/impl/activitystreams/property_endtime"
	propertyfirst "github.com/go-fed/activity/streams/impl/activitystreams/property_first"
	propertyfollowers "github.com/go-fed/activity/streams/impl/activitystreams/property_followers"
	propertyfollowing "github.com/go-fed/activity/streams/impl/activitystreams/property_following"
	propertyformertype "github.com/go-fed/activity/streams/impl/activitystreams/property_formertype"
	propertygenerator "github.com/go-fed/activity/streams/impl/activitystreams/property_generator"
	propertyheight "github.com/go-fed/activity/streams/impl/activitystreams/property_height"
	propertyhref "github.com/go-fed/activity/streams/impl/activitystreams/property_href"
	propertyhreflang "github.com/go-fed/activity/streams/impl/activitystreams/property_hreflang"
	propertyicon "github.com/go-fed/activity/streams/impl/activitystreams/property_icon"
	propertyid "github.com/go-fed/activity/streams/impl/activitystreams/property_id"
	propertyimage "github.com/go-fed/activity/streams/impl/activitystreams/property_image"
	propertyinbox "github.com/go-fed/activity/streams/impl/activitystreams/property_inbox"
	propertyinreplyto "github.com/go-fed/activity/streams/impl/activitystreams/property_inreplyto"
	propertyinstrument "github.com/go-fed/activity/streams/impl/activitystreams/property_instrument"
	propertyitems "github.com/go-fed/activity/streams/impl/activitystreams/property_items"
	propertylast "github.com/go-fed/activity/streams/impl/activitystreams/property_last"
	propertylatitude "github.com/go-fed/activity/streams/impl/activitystreams/property_latitude"
	propertyliked "github.com/go-fed/activity/streams/impl/activitystreams/property_liked"
	propertylikes "github.com/go-fed/activity/streams/impl/activitystreams/property_likes"
	propertylocation "github.com/go-fed/activity/streams/impl/activitystreams/property_location"
	propertylongitude "github.com/go-fed/activity/streams/impl/activitystreams/property_longitude"
	propertymediatype "github.com/go-fed/activity/streams/impl/activitystreams/property_mediatype"
	propertyname "github.com/go-fed/activity/streams/impl/activitystreams/property_name"
	propertynext "github.com/go-fed/activity/streams/impl/activitystreams/property_next"
	propertyobject "github.com/go-fed/activity/streams/impl/activitystreams/property_object"
	propertyoneof "github.com/go-fed/activity/streams/impl/activitystreams/property_oneof"
	propertyordereditems "github.com/go-fed/activity/streams/impl/activitystreams/property_ordereditems"
	propertyorigin "github.com/go-fed/activity/streams/impl/activitystreams/property_origin"
	propertyoutbox "github.com/go-fed/activity/streams/impl/activitystreams/property_outbox"
	propertypartof "github.com/go-fed/activity/streams/impl/activitystreams/property_partof"
	propertypreferredusername "github.com/go-fed/activity/streams/impl/activitystreams/property_preferredusername"
	propertyprev "github.com/go-fed/activity/streams/impl/activitystreams/property_prev"
	propertypreview "github.com/go-fed/activity/streams/impl/activitystreams/property_preview"
	propertypublished "github.com/go-fed/activity/streams/impl/activitystreams/property_published"
	propertyradius "github.com/go-fed/activity/streams/impl/activitystreams/property_radius"
	propertyrel "github.com/go-fed/activity/streams/impl/activitystreams/property_rel"
	propertyrelationship "github.com/go-fed/activity/streams/impl/activitystreams/property_relationship"
	propertyreplies "github.com/go-fed/activity/streams/impl/activitystreams/property_replies"
	propertyresult "github.com/go-fed/activity/streams/impl/activitystreams/property_result"
	propertyshares "github.com/go-fed/activity/streams/impl/activitystreams/property_shares"
	propertystartindex "github.com/go-fed/activity/streams/impl/activitystreams/property_startindex"
	propertystarttime "github.com/go-fed/activity/streams/impl/activitystreams/property_starttime"
	propertystreams "github.com/go-fed/activity/streams/impl/activitystreams/property_streams"
	propertysubject "github.com/go-fed/activity/streams/impl/activitystreams/property_subject"
	propertysummary "github.com/go-fed/activity/streams/impl/activitystreams/property_summary"
	propertytag "github.com/go-fed/activity/streams/impl/activitystreams/property_tag"
	propertytarget "github.com/go-fed/activity/streams/impl/activitystreams/property_target"
	propertyto "github.com/go-fed/activity/streams/impl/activitystreams/property_to"
	propertytotalitems "github.com/go-fed/activity/streams/impl/activitystreams/property_totalitems"
	propertytype "github.com/go-fed/activity/streams/impl/activitystreams/property_type"
	propertyunits "github.com/go-fed/activity/streams/impl/activitystreams/property_units"
	propertyupdated "github.com/go-fed/activity/streams/impl/activitystreams/property_updated"
	propertyurl "github.com/go-fed/activity/streams/impl/activitystreams/property_url"
	propertywidth "github.com/go-fed/activity/streams/impl/activitystreams/property_width"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// NewActivityStreamsActivityStreamsAccuracyProperty creates a new
// ActivityStreamsAccuracyProperty
func NewActivityStreamsAccuracyProperty() vocab.ActivityStreamsAccuracyProperty {
	return propertyaccuracy.NewActivityStreamsAccuracyProperty()
}

// NewActivityStreamsActivityStreamsActorProperty creates a new
// ActivityStreamsActorProperty
func NewActivityStreamsActorProperty() vocab.ActivityStreamsActorProperty {
	return propertyactor.NewActivityStreamsActorProperty()
}

// NewActivityStreamsActivityStreamsAltitudeProperty creates a new
// ActivityStreamsAltitudeProperty
func NewActivityStreamsAltitudeProperty() vocab.ActivityStreamsAltitudeProperty {
	return propertyaltitude.NewActivityStreamsAltitudeProperty()
}

// NewActivityStreamsActivityStreamsAnyOfProperty creates a new
// ActivityStreamsAnyOfProperty
func NewActivityStreamsAnyOfProperty() vocab.ActivityStreamsAnyOfProperty {
	return propertyanyof.NewActivityStreamsAnyOfProperty()
}

// NewActivityStreamsActivityStreamsAttachmentProperty creates a new
// ActivityStreamsAttachmentProperty
func NewActivityStreamsAttachmentProperty() vocab.ActivityStreamsAttachmentProperty {
	return propertyattachment.NewActivityStreamsAttachmentProperty()
}

// NewActivityStreamsActivityStreamsAttributedToProperty creates a new
// ActivityStreamsAttributedToProperty
func NewActivityStreamsAttributedToProperty() vocab.ActivityStreamsAttributedToProperty {
	return propertyattributedto.NewActivityStreamsAttributedToProperty()
}

// NewActivityStreamsActivityStreamsAudienceProperty creates a new
// ActivityStreamsAudienceProperty
func NewActivityStreamsAudienceProperty() vocab.ActivityStreamsAudienceProperty {
	return propertyaudience.NewActivityStreamsAudienceProperty()
}

// NewActivityStreamsActivityStreamsBccProperty creates a new
// ActivityStreamsBccProperty
func NewActivityStreamsBccProperty() vocab.ActivityStreamsBccProperty {
	return propertybcc.NewActivityStreamsBccProperty()
}

// NewActivityStreamsActivityStreamsBtoProperty creates a new
// ActivityStreamsBtoProperty
func NewActivityStreamsBtoProperty() vocab.ActivityStreamsBtoProperty {
	return propertybto.NewActivityStreamsBtoProperty()
}

// NewActivityStreamsActivityStreamsCcProperty creates a new
// ActivityStreamsCcProperty
func NewActivityStreamsCcProperty() vocab.ActivityStreamsCcProperty {
	return propertycc.NewActivityStreamsCcProperty()
}

// NewActivityStreamsActivityStreamsClosedProperty creates a new
// ActivityStreamsClosedProperty
func NewActivityStreamsClosedProperty() vocab.ActivityStreamsClosedProperty {
	return propertyclosed.NewActivityStreamsClosedProperty()
}

// NewActivityStreamsActivityStreamsContentProperty creates a new
// ActivityStreamsContentProperty
func NewActivityStreamsContentProperty() vocab.ActivityStreamsContentProperty {
	return propertycontent.NewActivityStreamsContentProperty()
}

// NewActivityStreamsActivityStreamsContextProperty creates a new
// ActivityStreamsContextProperty
func NewActivityStreamsContextProperty() vocab.ActivityStreamsContextProperty {
	return propertycontext.NewActivityStreamsContextProperty()
}

// NewActivityStreamsActivityStreamsCurrentProperty creates a new
// ActivityStreamsCurrentProperty
func NewActivityStreamsCurrentProperty() vocab.ActivityStreamsCurrentProperty {
	return propertycurrent.NewActivityStreamsCurrentProperty()
}

// NewActivityStreamsActivityStreamsDeletedProperty creates a new
// ActivityStreamsDeletedProperty
func NewActivityStreamsDeletedProperty() vocab.ActivityStreamsDeletedProperty {
	return propertydeleted.NewActivityStreamsDeletedProperty()
}

// NewActivityStreamsActivityStreamsDescribesProperty creates a new
// ActivityStreamsDescribesProperty
func NewActivityStreamsDescribesProperty() vocab.ActivityStreamsDescribesProperty {
	return propertydescribes.NewActivityStreamsDescribesProperty()
}

// NewActivityStreamsActivityStreamsDurationProperty creates a new
// ActivityStreamsDurationProperty
func NewActivityStreamsDurationProperty() vocab.ActivityStreamsDurationProperty {
	return propertyduration.NewActivityStreamsDurationProperty()
}

// NewActivityStreamsActivityStreamsEndTimeProperty creates a new
// ActivityStreamsEndTimeProperty
func NewActivityStreamsEndTimeProperty() vocab.ActivityStreamsEndTimeProperty {
	return propertyendtime.NewActivityStreamsEndTimeProperty()
}

// NewActivityStreamsActivityStreamsFirstProperty creates a new
// ActivityStreamsFirstProperty
func NewActivityStreamsFirstProperty() vocab.ActivityStreamsFirstProperty {
	return propertyfirst.NewActivityStreamsFirstProperty()
}

// NewActivityStreamsActivityStreamsFollowersProperty creates a new
// ActivityStreamsFollowersProperty
func NewActivityStreamsFollowersProperty() vocab.ActivityStreamsFollowersProperty {
	return propertyfollowers.NewActivityStreamsFollowersProperty()
}

// NewActivityStreamsActivityStreamsFollowingProperty creates a new
// ActivityStreamsFollowingProperty
func NewActivityStreamsFollowingProperty() vocab.ActivityStreamsFollowingProperty {
	return propertyfollowing.NewActivityStreamsFollowingProperty()
}

// NewActivityStreamsActivityStreamsFormerTypeProperty creates a new
// ActivityStreamsFormerTypeProperty
func NewActivityStreamsFormerTypeProperty() vocab.ActivityStreamsFormerTypeProperty {
	return propertyformertype.NewActivityStreamsFormerTypeProperty()
}

// NewActivityStreamsActivityStreamsGeneratorProperty creates a new
// ActivityStreamsGeneratorProperty
func NewActivityStreamsGeneratorProperty() vocab.ActivityStreamsGeneratorProperty {
	return propertygenerator.NewActivityStreamsGeneratorProperty()
}

// NewActivityStreamsActivityStreamsHeightProperty creates a new
// ActivityStreamsHeightProperty
func NewActivityStreamsHeightProperty() vocab.ActivityStreamsHeightProperty {
	return propertyheight.NewActivityStreamsHeightProperty()
}

// NewActivityStreamsActivityStreamsHrefProperty creates a new
// ActivityStreamsHrefProperty
func NewActivityStreamsHrefProperty() vocab.ActivityStreamsHrefProperty {
	return propertyhref.NewActivityStreamsHrefProperty()
}

// NewActivityStreamsActivityStreamsHreflangProperty creates a new
// ActivityStreamsHreflangProperty
func NewActivityStreamsHreflangProperty() vocab.ActivityStreamsHreflangProperty {
	return propertyhreflang.NewActivityStreamsHreflangProperty()
}

// NewActivityStreamsActivityStreamsIconProperty creates a new
// ActivityStreamsIconProperty
func NewActivityStreamsIconProperty() vocab.ActivityStreamsIconProperty {
	return propertyicon.NewActivityStreamsIconProperty()
}

// NewActivityStreamsActivityStreamsIdProperty creates a new
// ActivityStreamsIdProperty
func NewActivityStreamsIdProperty() vocab.ActivityStreamsIdProperty {
	return propertyid.NewActivityStreamsIdProperty()
}

// NewActivityStreamsActivityStreamsImageProperty creates a new
// ActivityStreamsImageProperty
func NewActivityStreamsImageProperty() vocab.ActivityStreamsImageProperty {
	return propertyimage.NewActivityStreamsImageProperty()
}

// NewActivityStreamsActivityStreamsInReplyToProperty creates a new
// ActivityStreamsInReplyToProperty
func NewActivityStreamsInReplyToProperty() vocab.ActivityStreamsInReplyToProperty {
	return propertyinreplyto.NewActivityStreamsInReplyToProperty()
}

// NewActivityStreamsActivityStreamsInboxProperty creates a new
// ActivityStreamsInboxProperty
func NewActivityStreamsInboxProperty() vocab.ActivityStreamsInboxProperty {
	return propertyinbox.NewActivityStreamsInboxProperty()
}

// NewActivityStreamsActivityStreamsInstrumentProperty creates a new
// ActivityStreamsInstrumentProperty
func NewActivityStreamsInstrumentProperty() vocab.ActivityStreamsInstrumentProperty {
	return propertyinstrument.NewActivityStreamsInstrumentProperty()
}

// NewActivityStreamsActivityStreamsItemsProperty creates a new
// ActivityStreamsItemsProperty
func NewActivityStreamsItemsProperty() vocab.ActivityStreamsItemsProperty {
	return propertyitems.NewActivityStreamsItemsProperty()
}

// NewActivityStreamsActivityStreamsLastProperty creates a new
// ActivityStreamsLastProperty
func NewActivityStreamsLastProperty() vocab.ActivityStreamsLastProperty {
	return propertylast.NewActivityStreamsLastProperty()
}

// NewActivityStreamsActivityStreamsLatitudeProperty creates a new
// ActivityStreamsLatitudeProperty
func NewActivityStreamsLatitudeProperty() vocab.ActivityStreamsLatitudeProperty {
	return propertylatitude.NewActivityStreamsLatitudeProperty()
}

// NewActivityStreamsActivityStreamsLikedProperty creates a new
// ActivityStreamsLikedProperty
func NewActivityStreamsLikedProperty() vocab.ActivityStreamsLikedProperty {
	return propertyliked.NewActivityStreamsLikedProperty()
}

// NewActivityStreamsActivityStreamsLikesProperty creates a new
// ActivityStreamsLikesProperty
func NewActivityStreamsLikesProperty() vocab.ActivityStreamsLikesProperty {
	return propertylikes.NewActivityStreamsLikesProperty()
}

// NewActivityStreamsActivityStreamsLocationProperty creates a new
// ActivityStreamsLocationProperty
func NewActivityStreamsLocationProperty() vocab.ActivityStreamsLocationProperty {
	return propertylocation.NewActivityStreamsLocationProperty()
}

// NewActivityStreamsActivityStreamsLongitudeProperty creates a new
// ActivityStreamsLongitudeProperty
func NewActivityStreamsLongitudeProperty() vocab.ActivityStreamsLongitudeProperty {
	return propertylongitude.NewActivityStreamsLongitudeProperty()
}

// NewActivityStreamsActivityStreamsMediaTypeProperty creates a new
// ActivityStreamsMediaTypeProperty
func NewActivityStreamsMediaTypeProperty() vocab.ActivityStreamsMediaTypeProperty {
	return propertymediatype.NewActivityStreamsMediaTypeProperty()
}

// NewActivityStreamsActivityStreamsNameProperty creates a new
// ActivityStreamsNameProperty
func NewActivityStreamsNameProperty() vocab.ActivityStreamsNameProperty {
	return propertyname.NewActivityStreamsNameProperty()
}

// NewActivityStreamsActivityStreamsNextProperty creates a new
// ActivityStreamsNextProperty
func NewActivityStreamsNextProperty() vocab.ActivityStreamsNextProperty {
	return propertynext.NewActivityStreamsNextProperty()
}

// NewActivityStreamsActivityStreamsObjectProperty creates a new
// ActivityStreamsObjectProperty
func NewActivityStreamsObjectProperty() vocab.ActivityStreamsObjectProperty {
	return propertyobject.NewActivityStreamsObjectProperty()
}

// NewActivityStreamsActivityStreamsOneOfProperty creates a new
// ActivityStreamsOneOfProperty
func NewActivityStreamsOneOfProperty() vocab.ActivityStreamsOneOfProperty {
	return propertyoneof.NewActivityStreamsOneOfProperty()
}

// NewActivityStreamsActivityStreamsOrderedItemsProperty creates a new
// ActivityStreamsOrderedItemsProperty
func NewActivityStreamsOrderedItemsProperty() vocab.ActivityStreamsOrderedItemsProperty {
	return propertyordereditems.NewActivityStreamsOrderedItemsProperty()
}

// NewActivityStreamsActivityStreamsOriginProperty creates a new
// ActivityStreamsOriginProperty
func NewActivityStreamsOriginProperty() vocab.ActivityStreamsOriginProperty {
	return propertyorigin.NewActivityStreamsOriginProperty()
}

// NewActivityStreamsActivityStreamsOutboxProperty creates a new
// ActivityStreamsOutboxProperty
func NewActivityStreamsOutboxProperty() vocab.ActivityStreamsOutboxProperty {
	return propertyoutbox.NewActivityStreamsOutboxProperty()
}

// NewActivityStreamsActivityStreamsPartOfProperty creates a new
// ActivityStreamsPartOfProperty
func NewActivityStreamsPartOfProperty() vocab.ActivityStreamsPartOfProperty {
	return propertypartof.NewActivityStreamsPartOfProperty()
}

// NewActivityStreamsActivityStreamsPreferredUsernameProperty creates a new
// ActivityStreamsPreferredUsernameProperty
func NewActivityStreamsPreferredUsernameProperty() vocab.ActivityStreamsPreferredUsernameProperty {
	return propertypreferredusername.NewActivityStreamsPreferredUsernameProperty()
}

// NewActivityStreamsActivityStreamsPrevProperty creates a new
// ActivityStreamsPrevProperty
func NewActivityStreamsPrevProperty() vocab.ActivityStreamsPrevProperty {
	return propertyprev.NewActivityStreamsPrevProperty()
}

// NewActivityStreamsActivityStreamsPreviewProperty creates a new
// ActivityStreamsPreviewProperty
func NewActivityStreamsPreviewProperty() vocab.ActivityStreamsPreviewProperty {
	return propertypreview.NewActivityStreamsPreviewProperty()
}

// NewActivityStreamsActivityStreamsPublishedProperty creates a new
// ActivityStreamsPublishedProperty
func NewActivityStreamsPublishedProperty() vocab.ActivityStreamsPublishedProperty {
	return propertypublished.NewActivityStreamsPublishedProperty()
}

// NewActivityStreamsActivityStreamsRadiusProperty creates a new
// ActivityStreamsRadiusProperty
func NewActivityStreamsRadiusProperty() vocab.ActivityStreamsRadiusProperty {
	return propertyradius.NewActivityStreamsRadiusProperty()
}

// NewActivityStreamsActivityStreamsRelProperty creates a new
// ActivityStreamsRelProperty
func NewActivityStreamsRelProperty() vocab.ActivityStreamsRelProperty {
	return propertyrel.NewActivityStreamsRelProperty()
}

// NewActivityStreamsActivityStreamsRelationshipProperty creates a new
// ActivityStreamsRelationshipProperty
func NewActivityStreamsRelationshipProperty() vocab.ActivityStreamsRelationshipProperty {
	return propertyrelationship.NewActivityStreamsRelationshipProperty()
}

// NewActivityStreamsActivityStreamsRepliesProperty creates a new
// ActivityStreamsRepliesProperty
func NewActivityStreamsRepliesProperty() vocab.ActivityStreamsRepliesProperty {
	return propertyreplies.NewActivityStreamsRepliesProperty()
}

// NewActivityStreamsActivityStreamsResultProperty creates a new
// ActivityStreamsResultProperty
func NewActivityStreamsResultProperty() vocab.ActivityStreamsResultProperty {
	return propertyresult.NewActivityStreamsResultProperty()
}

// NewActivityStreamsActivityStreamsSharesProperty creates a new
// ActivityStreamsSharesProperty
func NewActivityStreamsSharesProperty() vocab.ActivityStreamsSharesProperty {
	return propertyshares.NewActivityStreamsSharesProperty()
}

// NewActivityStreamsActivityStreamsStartIndexProperty creates a new
// ActivityStreamsStartIndexProperty
func NewActivityStreamsStartIndexProperty() vocab.ActivityStreamsStartIndexProperty {
	return propertystartindex.NewActivityStreamsStartIndexProperty()
}

// NewActivityStreamsActivityStreamsStartTimeProperty creates a new
// ActivityStreamsStartTimeProperty
func NewActivityStreamsStartTimeProperty() vocab.ActivityStreamsStartTimeProperty {
	return propertystarttime.NewActivityStreamsStartTimeProperty()
}

// NewActivityStreamsActivityStreamsStreamsProperty creates a new
// ActivityStreamsStreamsProperty
func NewActivityStreamsStreamsProperty() vocab.ActivityStreamsStreamsProperty {
	return propertystreams.NewActivityStreamsStreamsProperty()
}

// NewActivityStreamsActivityStreamsSubjectProperty creates a new
// ActivityStreamsSubjectProperty
func NewActivityStreamsSubjectProperty() vocab.ActivityStreamsSubjectProperty {
	return propertysubject.NewActivityStreamsSubjectProperty()
}

// NewActivityStreamsActivityStreamsSummaryProperty creates a new
// ActivityStreamsSummaryProperty
func NewActivityStreamsSummaryProperty() vocab.ActivityStreamsSummaryProperty {
	return propertysummary.NewActivityStreamsSummaryProperty()
}

// NewActivityStreamsActivityStreamsTagProperty creates a new
// ActivityStreamsTagProperty
func NewActivityStreamsTagProperty() vocab.ActivityStreamsTagProperty {
	return propertytag.NewActivityStreamsTagProperty()
}

// NewActivityStreamsActivityStreamsTargetProperty creates a new
// ActivityStreamsTargetProperty
func NewActivityStreamsTargetProperty() vocab.ActivityStreamsTargetProperty {
	return propertytarget.NewActivityStreamsTargetProperty()
}

// NewActivityStreamsActivityStreamsToProperty creates a new
// ActivityStreamsToProperty
func NewActivityStreamsToProperty() vocab.ActivityStreamsToProperty {
	return propertyto.NewActivityStreamsToProperty()
}

// NewActivityStreamsActivityStreamsTotalItemsProperty creates a new
// ActivityStreamsTotalItemsProperty
func NewActivityStreamsTotalItemsProperty() vocab.ActivityStreamsTotalItemsProperty {
	return propertytotalitems.NewActivityStreamsTotalItemsProperty()
}

// NewActivityStreamsActivityStreamsTypeProperty creates a new
// ActivityStreamsTypeProperty
func NewActivityStreamsTypeProperty() vocab.ActivityStreamsTypeProperty {
	return propertytype.NewActivityStreamsTypeProperty()
}

// NewActivityStreamsActivityStreamsUnitsProperty creates a new
// ActivityStreamsUnitsProperty
func NewActivityStreamsUnitsProperty() vocab.ActivityStreamsUnitsProperty {
	return propertyunits.NewActivityStreamsUnitsProperty()
}

// NewActivityStreamsActivityStreamsUpdatedProperty creates a new
// ActivityStreamsUpdatedProperty
func NewActivityStreamsUpdatedProperty() vocab.ActivityStreamsUpdatedProperty {
	return propertyupdated.NewActivityStreamsUpdatedProperty()
}

// NewActivityStreamsActivityStreamsUrlProperty creates a new
// ActivityStreamsUrlProperty
func NewActivityStreamsUrlProperty() vocab.ActivityStreamsUrlProperty {
	return propertyurl.NewActivityStreamsUrlProperty()
}

// NewActivityStreamsActivityStreamsWidthProperty creates a new
// ActivityStreamsWidthProperty
func NewActivityStreamsWidthProperty() vocab.ActivityStreamsWidthProperty {
	return propertywidth.NewActivityStreamsWidthProperty()
}