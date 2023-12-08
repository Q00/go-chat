package graphql

var Schema = `
schema {
	query: Query
	mutation: Mutation
	subscription: Subscription
}

type Query {
	# Get all rooms
	rooms(userId: String!): [Room!]!
	# Get Chat messages
	chatMessages(roomId: String!, limit: Int!, nextCursor: String): ChatMessagePage!
}

type Mutation {
	# add message to chat
	addChatMessage(message: ChatMessageInput!): Boolean!
	# create room
	createRoom(groupId: Int!, lectureId: Int!, users: [UserInput!]!): Room!
	# report Message
	reportChatMessage(chatMessageId: String!, roomId: String!, isReport: Boolean!): Boolean!
}

type Subscription {
	# subscribe to chat messages
	chatMessage(roomId: String!): ChatMessage!
}

type Room {
	# Room ID
	id: ID!
	# Group ID
	groupId: Int!
	# Lecture ID
	lectureId: Int!
	# Users
	users: [User!]!
}

type User {
	# User ID
	id: ID!
	# User name
	name: String!
	# User Type
	userType: UserType!
}

input UserInput {
	# User ID
	id: String!
	# User name
	name: String!
	# User Type
	userType: UserType!
}

scalar Time

type ChatMessage {
	# Chat message ID
	id: ID!
	# Room ID
	roomId: String!
	# User
	user: User!
	# Message
	message: String!
	# ChatMessageType
	chatMessageType: ChatMessageType!
	createdAt: Time!
	isReport: Boolean
}

input ChatMessageInput {
	# Room ID
	roomId: String!
	# User
	user: UserInput!
	# Message
	message: String!
	# ChatMessageType
	chatMessageType: ChatMessageType!
}

enum ChatMessageType {
	MESSAGE
	EMOJI
	IMAGE
}

enum UserType {
	STUDENT
	TEACHER
	ADMIN
}

type ChatMessagePage {
	# Chat messages
	chatMessages: [ChatMessage!]!
	# Last key
	nextCursor: String
}
`
