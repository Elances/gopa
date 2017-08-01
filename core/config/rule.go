/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

// MatchRule is container of rules
type MatchRule struct {
	Contain    []string
	NotContain []string
	Prefix     []string
	Suffix     []string
}

// ShouldMatchRule means some rule should match
type ShouldMatchRule struct {
	*MatchRule
}

// MustMatchRule means some rule must match
type MustMatchRule struct {
	*MatchRule
}

// MustNotMatchRule means some rule must not match
type MustNotMatchRule struct {
	*MatchRule
}
