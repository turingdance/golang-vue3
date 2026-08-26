package tokenkit

type Realm struct {
	User     string
	AccId    string
	TenantId uint64
	Avatar   string
	NickName string
	RoleKey  string
	Extern   map[string]any
}

// RealmBuilder 建造器
type RealmBuilder struct {
	r Realm
}

func NewRealmBuilder() *RealmBuilder {
	return &RealmBuilder{
		r: Realm{
			Extern: make(map[string]any, 0),
		},
	}
}

func (b *RealmBuilder) User(v string) *RealmBuilder {
	b.r.User = v
	return b
}

func (b *RealmBuilder) AccId(v string) *RealmBuilder {
	b.r.AccId = v
	return b
}

func (b *RealmBuilder) TenantId(v uint64) *RealmBuilder {
	b.r.TenantId = v
	return b
}

func (b *RealmBuilder) RoleKey(v string) *RealmBuilder {
	b.r.RoleKey = v
	return b
}

func (b *RealmBuilder) Avatar(v string) *RealmBuilder {
	b.r.Avatar = v
	return b
}

func (b *RealmBuilder) NickName(v string) *RealmBuilder {
	b.r.NickName = v
	return b
}

func (b *RealmBuilder) Append(data map[string]any) *RealmBuilder {
	return b
}

// Build 生成最终对象
func (b *RealmBuilder) Build() Realm {
	return b.r
}

// Build 生成最终对象
func (b *RealmBuilder) BuildPtr() *Realm {
	return &b.r
}
