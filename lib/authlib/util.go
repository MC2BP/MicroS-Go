package authlib

func (r Roles) Has(role string) bool {
	for _, userRole := range r {
		if userRole == role {
			return true
		}
	}
	return false
}

func (o ObjectPermission) HasRoleOn(object, role string) bool {
	roles, ok := o[object]
	if !ok {
		return false
	}
	return roles.Has(role)
}

func (p Permission) HasRoleOnOjbectType(objectType, object, role string) bool {
	objectPermission, ok := p[objectType]
	if !ok {
		return false
	}
	return objectPermission.HasRoleOn(object, role)
}
