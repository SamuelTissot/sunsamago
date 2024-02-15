package sunsama

import (
	"strings"

	"github.com/shurcooL/graphql"
)

type eventContext struct {
	name    string
	private bool
	channel string
}

type eventContexts []eventContext

func (c eventContexts) Context() string {
	names := make([]string, len(c))
	for i, ec := range c {
		names[i] = ec.name
	}

	return strings.Join(names, ":")
}

func (c eventContexts) Channel() string {
	names := make([]string, len(c))
	for i, ec := range c {
		names[i] = ec.channel
	}

	return strings.Join(names, ":")
}

func (c eventContexts) BelongsTo(names ...string) bool {
	for _, name := range names {
		name = strings.ToLower(name)

		for _, ec := range c {
			if strings.ToLower(ec.channel) == name {
				return true
			}

			if strings.ToLower(ec.name) == name {
				return true
			}
		}
	}

	return false
}

func (c eventContexts) IsPrivate() bool {
	for _, ec := range c {
		if ec.private {
			return true
		}
	}

	return false
}

type stream struct {
	ID               graphql.String `graphql:"_id"`
	StreamName       graphql.String
	Description      graphql.String
	Category         graphql.Boolean
	Personal         graphql.Boolean
	CategoryStreamID graphql.String `graphql:"categoryStreamId"`
}

type streams []stream

func (streams streams) contextsForIds(ids []graphql.String) []eventContext {
	contexts := make([]eventContext, len(ids))
	tree := streams.tree()

	for i, id := range ids {
		for parent, childrens := range tree {
			if parent.ID == id {
				contexts[i] = eventContext{
					name:    string(parent.StreamName),
					private: bool(parent.Personal),
					channel: string(parent.StreamName),
				}
				continue
			}

			for _, child := range childrens {
				if child.ID == id {
					contexts[i] = eventContext{
						name:    string(parent.StreamName),
						private: bool(child.Personal),
						channel: string(child.StreamName),
					}

					continue
				}
			}

		}
	}

	return contexts
}

func (streams streams) tree() map[stream][]stream {
	tree := make(map[stream][]stream)
	for _, s := range streams {
		if s.CategoryStreamID != "" {
			parent := stream{}
			for _, p := range streams {
				if p.ID == s.CategoryStreamID {
					parent = p
					break
				}
			}
			tree[parent] = append(tree[parent], s)
		}
	}
	return tree
}
