package handlers

import (
	"github.com/gofiber/fiber/v2"
	"main/utils"
	"main/discord"
)

func ProfileHandler(c *fiber.Ctx) error {
	err := utils.Discord.GetUser("916091258013892699")
	if err != nil {
		return err
	}

	return c.Render("terrifictable", fiber.Map{
		"discord_user": FormatUser(utils.Discord.User),
		"badges": FlagToBadges(utils.Discord.User.PublicFlags),
	})
}


func FormatUser(user discord.User) discord.User {
	return user
}

func FlagToBadges(flags int64) []string {
	var badges []string

	if flags & discord.FUser_staff == discord.FUser_staff {
		badges = append(badges, "Staff")
	}
	if flags & discord.FUser_partner == discord.FUser_partner {
		badges = append(badges, "Partner")
	}
	if flags & discord.FUser_hypesquad == discord.FUser_hypesquad {
		badges = append(badges, "Hypesquad")
	}
	if flags & discord.FUser_bughunter_1 == discord.FUser_bughunter_1 {
		badges = append(badges, "Bughunter lvl 1")
	}
	if flags & discord.FUser_hypesquad_brilliance == discord.FUser_hypesquad_brilliance {
		badges = append(badges, "Hypesquad Brilliance")
	}
	if flags & discord.FUser_hypesquad_balance == discord.FUser_hypesquad_balance {
		badges = append(badges, "Hypesquad Balance")
	}
	if flags & discord.FUser_hypesquad_bravery == discord.FUser_hypesquad_bravery {
		badges = append(badges, "Hypesquad Bravery")
	}
	if flags & discord.FUser_premium_eary_supporter == discord.FUser_premium_eary_supporter {
		badges = append(badges, "Early Supporter")
	}
	if flags & discord.FUser_team_pseudo_user == discord.FUser_team_pseudo_user {
		badges = append(badges, "Team of users")
	}
	if flags & discord.FUser_bughunter_2 == discord.FUser_bughunter_2 {
		badges = append(badges, "Bughunter lvl 2")
	}
	if flags & discord.FUser_verified_bot == discord.FUser_verified_bot {
		badges = append(badges, "Verified Bot")
	}
	if flags & discord.FUser_verified_developer == discord.FUser_verified_developer {
		badges = append(badges, "verified bot developer")
	}
	if flags & discord.FUser_certified_moderator == discord.FUser_certified_moderator {
		badges = append(badges, "Moderator (take a shower)")
	}
	if flags & discord.FUser_bot_http_interaction == discord.FUser_bot_http_interaction {
		badges = append(badges, "HTTP interaction Bot")
	}
	if flags & discord.FUser_active_developer == discord.FUser_active_developer {
		badges = append(badges, "Active bot developer")
	}
	return badges
}
