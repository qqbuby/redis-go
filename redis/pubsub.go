// The MIT License (MIT)
//
// Copyright (c) 2016 Roy Xu

package redis

type respPubSub interface {
	// PSUBSCRIBE pattern [pattern ...]
	// Listen for messages published to channels matching the given patterns

	// PUBSUB subcommand [argument [argument ...]]
	// Inspect the state of the Pub/Sub subsystem

	// PUBSUB CHANNELS [pattern]
	// Lists the currently active channels.
	// Array reply: a list of active channels, optionally matching the specified pattern.

	// PUBSUB NUMSUB [channel-1 ... channel-N]
	// Returns the number of subscribers (not counting clients subscribed to patterns) for the specified channels.
	// Array reply:
	//     a list of channels and number of subscribers for every channel.
	//     The format is channel, count, channel, count, ..., so the list is flat.
	//     The order in which the channels are listed is the same as the order of the channels specified in the command call.
	// Note that it is valid to call this command without channels. In this case it will just return an empty list.

	// PUBSUB NUMPAT
	// Returns the number of subscriptions to patterns (that are performed using the PSUBSCRIBE command).
	// Note that this is not just the count of clients subscribed to patterns but the total number of patterns all the clients are subscribed to.
	// Integer reply: the number of patterns all the clients are subscribed to.

	// PUBLISH channel message
	// Post a message to a channel
	// Integer reply: the number of clients that received the message.

	// PUNSUBSCRIBE [pattern [pattern ...]]
	// Stop listening for messages posted to channels matching the given patterns

	// SUBSCRIBE channel [channel ...]
	// Listen for messages published to the given channels

	// UNSUBSCRIBE [channel [channel ...]]
	// Stop listening for messages posted to the given channels
}
