import telebot
import json
from datetime import datetime

bot = telebot.TeleBot("token")

weekday_order = ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY"]

@bot.message_handler(commands=['mayaToday'])
def today_menu(message):
    with open('menu.json', 'r') as file:
        menu = json.load(file)
    today = datetime.now().strftime("%A").upper()
    if today in menu:
        item1 = menu[today]["Item 1"]
        item2 = menu[today]["Item 2"]
        response = f"{today}\n\n1. {item1}\n\n2. {item2}"
    else:
        response = "Sorry, today's menu is not available."
    bot.reply_to(message, response)

@bot.message_handler(commands=['mayaWeek'])
def week_menu(message):
    with open('menu.json', 'r') as file:
        menu = json.load(file)
    sorted_days = sorted(menu.keys(), key=lambda day: weekday_order.index(day.upper()))
    response = ""
    for day in sorted_days:
        items = menu[day]
        response += f"{day}\n1. {items['Item 1']}\n2. {items['Item 2']}\n\n"
    bot.reply_to(message, response)

bot.polling()
