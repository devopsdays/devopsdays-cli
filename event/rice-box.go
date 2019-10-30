package event

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `conduct.md.tmpl`,
		FileModTime: time.Unix(1521888192, 0),
		Content:     string("+++\nTitle = \"Conduct\"\nType = \"event\"\nDescription = \"Code of conduct for devopsdays [[ .City ]] [[ .Year ]]\"\n+++\n\n## ANTI-HARASSMENT POLICY\n\nDevopsdays is dedicated to providing a harassment-free conference experience for everyone, regardless of gender, sexual orientation, disability, physical appearance, body size, race, or religion. We do not tolerate harassment of conference participants in any form. Sexual language and imagery is not appropriate for any conference venue, including talks. Conference participants violating these rules may be sanctioned or expelled from the conference without a refund at the discretion of the conference organizers.\n\nHarassment includes offensive verbal comments related to gender, sexual orientation, disability, physical appearance, body size, race, religion, sexual images in public spaces, deliberate intimidation, stalking, following, harassing photography or recording, sustained disruption of talks or other events, inappropriate physical contact, and unwelcome sexual attention. Participants asked to stop any harassing behavior are expected to comply immediately.\n\nExhibitors in the expo hall, sponsor or vendor booths, or similar activities are also subject to the anti-harassment policy. In particular, exhibitors should not use sexualized images, activities, or other material. Booth staff (including volunteers) should not use sexualized clothing/uniforms/costumes, or otherwise create a sexualized environment.\n\nIf a participant engages in harassing behavior, the conference organizers may take any action they deem appropriate, including warning the offender or expulsion from the conference with no refund.\n\nIf you are being harassed, notice that someone else is being harassed, or have any other concerns, please contact a member of conference staff immediately.\n\nConference staff can be identified by distinct staff badges. Conference staff will be happy to help participants contact hotel/venue security or local law enforcement, provide escorts, or otherwise assist those experiencing harassment to feel safe for the duration of the conference. We value your attendance.\n\nWe expect participants to adhere to the code of conduct at all conference venues and conference-related social events.\n\n## CODE OF CONDUCT\n\nI. I am an attendee at devopsdays, learning from and sharing with other devopsdays attendees in an effort to better myself and my industry. I co-create the experience with fellow attendees. I am prepared to give my energy, presence and sensitivity to creating the best possible experience for myself and others.\n\nII. I am coming to devopsdays to interact with people. I understand that imagery and language which is suggestive or derogatory will offend and make people uncomfortable. I also understand that people may have boundaries and sensibilities different from my own. I will accept without question when informed that something is offensive or unacceptable in the context of the devopsdays event.\n\nIII. I will never intentionally harass or offend another attendee regardless of gender, sexual orientation, disability, appearance, size, race or religion and will not abide another attendee being harassed or offended. If I am aware that anyone is uncomfortable or unsafe, I will notify those giving offense and the devopsdays event organizers.\n\nIV. If I am offended or harassed, I will inform people around me who make me feel safe and the event organizers. If I feel safe, at my discretion, I will inform those giving offense of the specific actions with the hope that the other party is well-intentioned and ignorant, but I am under no obligation to do so.\n\nV. I understand that people are different and I attempt to be forgiving of others actions at the level of their sincere intent, but my priority is protecting my safety and the safety of others. I will act without hesitation or reservation until there are no question of the safety of all parties.\n\nVI. I trust the devopsdays organizers and attendees will co-create the best possible experience for everyone involved, as I will. I believe devopsdays is about empowering people and I will not forget I am empowered to create a safe and nurturing environment. If I or any other attendee violates this aspect of the event, I expect the conference organizers to protect the attendees by direct action, including expelling those in violation and contacting the proper authorities.\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `contact.md.tmpl`,
		FileModTime: time.Unix(1521888192, 0),
		Content:     string("+++\nTitle = \"Contact\"\nType = \"event\"\nDescription = \"Contact information for devopsdays [[ .City ]] [[ .Year ]]\"\n+++\n\nIf you'd like to contact us by email: {{< email_organizers >}}\n\n**Our local team**\n\n{{< list_organizers >}}\n\n**The core devopsdays organizer group**\n\n{{< list_core >}}\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    `location.md.tmpl`,
		FileModTime: time.Unix(1521888192, 0),
		Content:     string("+++\nTitle = \"Location\"\nType = \"event\"\nDescription = \"Location for devopsdays [[ .City ]] [[ .Year ]]\"\n+++\n\nWatch this space for information about the venue including address, map/direction, parking/transit, and any hotel details.\n\n<!-- Uncomment this only if you have set the coordinates for your location in the config yaml. Get Latitude and Longitude of a Point: http://itouchmap.com/latlong.html -->\n<!-- {{< event_map >}} -->\n"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    `program.md.tmpl`,
		FileModTime: time.Unix(1521888192, 0),
		Content:     string("+++\nTitle = \"Program\"\nType = \"program\"\nDescription = \"Program for devopsdays [[ .City ]] [[ .Year ]]\"\n+++\n"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    `propose.md.tmpl`,
		FileModTime: time.Unix(1521888192, 0),
		Content:     string("+++\nTitle = \"Propose\"\nType = \"event\"\nDescription = \"Propose a talk for devopsdays [[ .City ]] [[ .Year ]]\"\n+++\n  {{< cfp_dates >}}\n\n<hr>\n\nThere are three ways to propose a topic at devopsdays:\n<ol>\n  <li><strong><em>A 30-minute talk</em></strong> presented during the conference, usually in the mornings.</li>\n  <li><strong><em>An Ignite talk</em></strong> presented during the <a href=\"/pages/ignite-talks-format\">Ignite sessions</a> (scheduling varies). These are 5 minutes slots with slides changing every 15 seconds (20 slides total).</li>\n  <li><strong><em>Open Space</em></strong>: If you'd like to lead a group discussion during the attendee-suggested <a href=\"/pages/open-space-format\">Open Space</a> breakout sessions, it is not necessary to propose it ahead of time. Those topics are suggested in person at the conference. If you'd like to demo your product or service, you should <a href=\"../sponsor\">sponsor the event</a> and demo it at your table.\n</ol>\n\n<hr>\n\nChoosing talks is part art, part science; here are some factors we consider when trying to assemble the best possible program for our local audience:\n\n- _broad appeal_: How will your talk play out in a room of people with a variety of backgrounds? Technical deep dives need more levels to provide value for the whole room, some of whom might not use your specific tool.\n- _new local presenters_: You are the only one who can tell your story. We are very interested in the challenges and successes being experienced in our local area. We are happy to provide guidance/coaching for new speakers upon request.\n- _under-represented voices_: We want to hear all voices, including those that may speak less frequently at similar events. Whether you're in a field not typically thought of as a technology field, you're in a large, traditional organization, or you're the only person at your organization with your background, we are interested in your unique experience.\n- _original content_: We will consider talks that have already been presented elsewhere, but we prefer talks that the local area isn't likely to have already seen.\n- _no third-party submissions_: This is a small community-driven event, and speakers need to be directly engaged with the organizers and attendees. If a PR firm or your marketing department is proposing the talk, you've already shown that as a speaker you're distant from the process.\n- _no vendor pitches_: As much as we value vendors and sponsors, we are not going to accept a talk that appears to be a pitch for your product.\n\n<hr>\n\n<strong>How to submit a proposal:</strong> Send an email to [{{< email_proposals >}}] with the following information\n<ol>\n\t<li>Type (presentation, panel discussion, ignite)</li>\n\t<li>Proposal Title (can be changed later)</li>\n\t<li>Description (several sentences explaining what attendees will learn)</li>\n</ol>\n"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    `registration.md.tmpl`,
		FileModTime: time.Unix(1521888192, 0),
		Content:     string("+++\nTitle = \"Registration\"\nType = \"event\"\nDescription = \"Registration for devopsdays [[ .City ]] [[ .Year ]]\"\n+++\n\n<div style=\"width:100%; text-align:left;\">\n\nEmbed registration iframe/link/etc.\n</div></div>\n</div>\n"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    `sponsor.md.tmpl`,
		FileModTime: time.Unix(1521888192, 0),
		Content:     string("+++\nTitle = \"Sponsor\"\nType = \"event\"\nDescription = \"Sponsor devopsdays [[ .City ]] [[ .Year ]]\"\n+++\n\nWe greatly value sponsors for this open event.  If you are interested in sponsoring, please drop us an email at [{{< email_organizers >}}].\n\n<hr>\n\ndevopsdays is a self-organizing conference for practitioners that depends on sponsorships. We do not have vendor booths, sell product presentations, or distribute attendee contact lists. Sponsors have the opportunity to have short elevator pitches during the program and will get recognition on the website and social media before, during and after the event. Sponsors are encouraged to represent themselves by actively participating and engaging with the attendees as peers. Any attendee also has the opportunity to demo products/projects as part of an open space session.\n<p>\nGold sponsors get a full table and Silver sponsors a shared table where they can interact with those interested to come visit during breaks. All attendees are welcome to propose any subject they want during the open spaces, but this is a community-focused conference, so heavy marketing will probably work against you when trying to make a good impression on the attendees.\n<p>\nThe best thing to do is send engineers to interact with the experts at devopsdays on their own terms.\n<p>\n\n<!--\n<hr/>\n\n<div style=\"width:590px\">\n<table border=1 cellspacing=1>\n  <tr>\n    <th><i>packages</i></th>\n    <th><center><b><u>Bronze<br />1000 usd</u></center></b></th>\n    <th><center><b><u>Silver<br />3000 usd</u></center></b></th>\n    <th><center><b><u>Gold<br />5000 usd</u></center></b></th>\n    <th></th>\n  </tr>\n<tr><td>2 included tickets</td><td bgcolor=\"gold\">&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td></tr>\n<tr><td>logo on event website</td><td bgcolor=\"gold\">&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td></tr>\n<tr><td>logo on shared slide, rotating during breaks</td><td bgcolor=\"gold\">&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td></tr>\n<tr><td>logo on all email communication</td><td>&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td></tr>\n<tr><td>logo on its own slide, rotating during breaks</td><td>&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td></tr>\n<tr><td>1 minute pitch to full audience (including streaming audience)</td><td>&nbsp;</td><td>&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td></tr></tr>\n<tr><td>2 additional tickets (4 in total)</td><td>&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td><td>&nbsp;</td></tr>\n<tr><td>4 additional tickets (6 in total)</td><td>&nbsp;</td><td>&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td></tr>\n<tr><td>shared table for swag</td><td>&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td><td>&nbsp;</td></tr>\n<tr><td>booth/table space</td><td>&nbsp;</td><td>&nbsp;</td><td bgcolor=\"gold\">&nbsp;</td></tr>\n</table>\n<hr/>\nThere are also opportunities for exclusive special sponsorships. We'll have sponsors for various events with special privileges for the sponsors of these events. If you are interested in special sponsorships or have a creative idea about how you can support the event, send us an email.\n<br/>\n<br/>\n\n<br>\n<br>\n<table border=1 cellspacing=1>\n  <tr>\n    <th><i>Sponsor FAQ</i></th>\n    <th><center><b>Answers to questions frequently asked by sponsors&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</center></b></th>\n    <th></th>\n  </tr>\n<tr><td>What dates/times can we set up and tear down?</td><td></td></tr>\n<tr><td>How do we ship to the venue?</td><td></td></tr>\n<tr><td>How do we ship from the venue?</td><td></td></tr>\n<tr><td>Whom should we send?</td><td></td></tr>\n<tr><td>What should we expect regarding electricity? (how much, any fees, etc)</td><td></td></tr>\n<tr><td>What should we expect regarding WiFi? (how much, any fees, etc)</td><td></td></tr>\n<tr><td>How do we order additional A/V equipment?</td><td></td></tr>\n<tr><td>Additional important details</td><td></td></tr>\n</table>\n</div>\n\n-->\n<hr/>\n"),
	}
	file9 := &embedded.EmbeddedFile{
		Filename:    `welcome.md.tmpl`,
		FileModTime: time.Unix(1521890242, 0),
		Content:     string("+++\nTitle = \"devopsdays [[ .City ]] [[ .Year ]]\"\nType = \"welcome\"\naliases = [\"/events/[[ .Slug ]]/\"]\nDescription = \"devopsdays [[ .City ]] [[ .Year ]]\"\n+++\n\n<!-- <div style=\"text-align:center;\">\n  {{< event_logo >}}\n</div> -->\n\n<div class = \"row\">\n  <div class = \"col-md-2\">\n    <strong>Dates</strong>\n  </div>\n  <div class = \"col-md-8\">\n    {{< event_start >}} - {{< event_end >}}\n  </div>\n</div>\n\n<!-- <div class = \"row\">\n  <div class = \"col-md-2\">\n    <strong>Location</strong>\n  </div>\n  <div class = \"col-md-8\">\n    {{< event_location >}}\n  </div>\n</div> -->\n\n<!-- <div class = \"row\">\n  <div class = \"col-md-2\">\n    <strong>Register</strong>\n  </div>\n  <div class = \"col-md-8\">\n    {{< event_link page=\"registration\" text=\"Register to attend the conference!\" >}}\n  </div>\n</div> -->\n\n<!-- <div class = \"row\">\n  <div class = \"col-md-2\">\n    <strong>Propose</strong>\n  </div>\n  <div class = \"col-md-8\">\n    {{< event_link page=\"propose\" text=\"Propose a talk!\" >}}\n  </div>\n</div> -->\n\n<!-- <div class = \"row\">\n  <div class = \"col-md-2\">\n    <strong>Program</strong>\n  </div>\n  <div class = \"col-md-8\">\n    View the {{< event_link page=\"program\" text=\"program.\" >}}\n  </div>\n</div> -->\n\n<!-- <div class = \"row\">\n  <div class = \"col-md-2\">\n    <strong>Speakers</strong>\n  </div>\n  <div class = \"col-md-8\">\n    Check out the {{< event_link page=\"speakers\" text=\"speakers!\" >}}\n  </div>\n</div> -->\n\n<div class = \"row\">\n  <div class = \"col-md-2\">\n    <strong>Sponsors</strong>\n  </div>\n  <div class = \"col-md-8\">\n    {{< event_link page=\"sponsor\" text=\"Sponsor the conference!\" >}}\n  </div>\n</div>\n\n<div class = \"row\">\n  <div class = \"col-md-2\">\n    <strong>Contact</strong>\n  </div>\n  <div class = \"col-md-8\">\n    {{< event_link page=\"contact\" text=\"Get in touch with the organizers\" >}}\n  </div>\n</div>\n\n<!-- Uncomment if you added your city twitter name -->\n<!--\n{{< event_twitter >}}\n-->\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1521890242, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // conduct.md.tmpl
			file3, // contact.md.tmpl
			file4, // location.md.tmpl
			file5, // program.md.tmpl
			file6, // propose.md.tmpl
			file7, // registration.md.tmpl
			file8, // sponsor.md.tmpl
			file9, // welcome.md.tmpl

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1521890242, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"conduct.md.tmpl":      file2,
			"contact.md.tmpl":      file3,
			"location.md.tmpl":     file4,
			"program.md.tmpl":      file5,
			"propose.md.tmpl":      file6,
			"registration.md.tmpl": file7,
			"sponsor.md.tmpl":      file8,
			"welcome.md.tmpl":      file9,
		},
	})
}
