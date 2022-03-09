/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/provider-azurerm-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ChannelAlexas returns a ChannelAlexaInformer.
	ChannelAlexas() ChannelAlexaInformer
	// ChannelDirectLineSpeeches returns a ChannelDirectLineSpeechInformer.
	ChannelDirectLineSpeeches() ChannelDirectLineSpeechInformer
	// ChannelDirectlines returns a ChannelDirectlineInformer.
	ChannelDirectlines() ChannelDirectlineInformer
	// ChannelEmails returns a ChannelEmailInformer.
	ChannelEmails() ChannelEmailInformer
	// ChannelFacebooks returns a ChannelFacebookInformer.
	ChannelFacebooks() ChannelFacebookInformer
	// ChannelLines returns a ChannelLineInformer.
	ChannelLines() ChannelLineInformer
	// ChannelMsTeamses returns a ChannelMsTeamsInformer.
	ChannelMsTeamses() ChannelMsTeamsInformer
	// ChannelSlacks returns a ChannelSlackInformer.
	ChannelSlacks() ChannelSlackInformer
	// ChannelSmses returns a ChannelSmsInformer.
	ChannelSmses() ChannelSmsInformer
	// ChannelWebChats returns a ChannelWebChatInformer.
	ChannelWebChats() ChannelWebChatInformer
	// ChannelsRegistrations returns a ChannelsRegistrationInformer.
	ChannelsRegistrations() ChannelsRegistrationInformer
	// Connections returns a ConnectionInformer.
	Connections() ConnectionInformer
	// ServiceAzureBots returns a ServiceAzureBotInformer.
	ServiceAzureBots() ServiceAzureBotInformer
	// WebApps returns a WebAppInformer.
	WebApps() WebAppInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ChannelAlexas returns a ChannelAlexaInformer.
func (v *version) ChannelAlexas() ChannelAlexaInformer {
	return &channelAlexaInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChannelDirectLineSpeeches returns a ChannelDirectLineSpeechInformer.
func (v *version) ChannelDirectLineSpeeches() ChannelDirectLineSpeechInformer {
	return &channelDirectLineSpeechInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChannelDirectlines returns a ChannelDirectlineInformer.
func (v *version) ChannelDirectlines() ChannelDirectlineInformer {
	return &channelDirectlineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChannelEmails returns a ChannelEmailInformer.
func (v *version) ChannelEmails() ChannelEmailInformer {
	return &channelEmailInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChannelFacebooks returns a ChannelFacebookInformer.
func (v *version) ChannelFacebooks() ChannelFacebookInformer {
	return &channelFacebookInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChannelLines returns a ChannelLineInformer.
func (v *version) ChannelLines() ChannelLineInformer {
	return &channelLineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChannelMsTeamses returns a ChannelMsTeamsInformer.
func (v *version) ChannelMsTeamses() ChannelMsTeamsInformer {
	return &channelMsTeamsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChannelSlacks returns a ChannelSlackInformer.
func (v *version) ChannelSlacks() ChannelSlackInformer {
	return &channelSlackInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChannelSmses returns a ChannelSmsInformer.
func (v *version) ChannelSmses() ChannelSmsInformer {
	return &channelSmsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChannelWebChats returns a ChannelWebChatInformer.
func (v *version) ChannelWebChats() ChannelWebChatInformer {
	return &channelWebChatInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ChannelsRegistrations returns a ChannelsRegistrationInformer.
func (v *version) ChannelsRegistrations() ChannelsRegistrationInformer {
	return &channelsRegistrationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Connections returns a ConnectionInformer.
func (v *version) Connections() ConnectionInformer {
	return &connectionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ServiceAzureBots returns a ServiceAzureBotInformer.
func (v *version) ServiceAzureBots() ServiceAzureBotInformer {
	return &serviceAzureBotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WebApps returns a WebAppInformer.
func (v *version) WebApps() WebAppInformer {
	return &webAppInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
