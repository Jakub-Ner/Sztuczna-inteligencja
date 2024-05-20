from experta import *

class Malakser(Fact):pass
class MalakserKomponent(Fact):pass
class PrzygotowanyMalakser(Fact):pass

class Skladnik(Fact):pass

class Problem(Fact):pass
class Przyczyna(Fact):pass
class Rozwaizanie(Fact):pass

class STRONA_OSTRZA:
    TNACA = "tnąca na talarki"
    TRACA = "trąca na wiórki"

class ZASILANIE:
    BRAK = "brak prądu"
    JEST = "jest prąd"

    

PROBLEM1 = "Urządzenie nie włącza się"
PROBLEM2 = "Urządzenie się nagrzewa po uruchomieniu"
PROBLEM3 = "Urządzenie śmierdzi spalenizną po uruchomieniu"

PRZYCZYNA1 = "Brak prądu"
PRZYCZYNA2 = "Włącznik jest zepsuty"
PRZYCZYNA3 = "Silnik jest uszkodzony"
PRZYCZYNA4 = "Składniki są zbyt twarde"

ROZWIAZANIE1 = "Podłącz urządzenie do sprawnego kontaktu"
ROZWIAZANIE2 = "Oddaj urządzenie do serwisu"
ROZWIAZANIE3 = "Nie umieszczaj twardych składników w pojemniku malaksera"




class MalakserExpert(KnowledgeEngine):
    @DefFacts()
    def komponenty_scalone(self):
        yield MalakserKomponent(nazwa="silnik") 
        yield MalakserKomponent(nazwa="kabel zasilajacy", stan=ZASILANIE.BRAK)
        yield MalakserKomponent(nazwa="włącznik", wlaczony=False)
        yield PrzygotowanyMalakser(nazwa="nie")

    @Rule(AND(
        Malakser(),
        PrzygotowanyMalakser(nazwa="tak"),
        MalakserKomponent(nazwa="kabel zasilajacy", stan=ZASILANIE.JEST),
        MalakserKomponent(nazwa="włącznik", wlaczony=True),
        MalakserKomponent(nazwa="ostrze dwustronne", strona=MATCH.strona),
        Skladnik(nazwa=MATCH.nazwa),
        OR(
            MalakserKomponent(nazwa="popychacz"),
            NOT(MalakserKomponent(nazwa="popychacz")),
        ),
    ))
    def uzyj_malaksera(self, nazwa, strona):
        with_popychacz = [fact["nazwa"] for fact in self.facts.values() if isinstance(fact, MalakserKomponent)]
        skladniki = ", ".join(nazwa)
        if "popychacz" in with_popychacz:
            print(f"Zawartość blendera: {skladniki}, starta na {strona.split(' ')[-1]} z użyciem popychacza")
        else:
            print(f"Zawartość blendera: {skladniki}, starta na surówkę")

    @Rule(
        Malakser(),
        MalakserKomponent(nazwa="pokrywa"),
        MalakserKomponent(nazwa="ostrze dwustronne"),
        MalakserKomponent(nazwa="mocowanie na ostrze"),
        MalakserKomponent(nazwa="pojemnik"),
        salience=1)
    def zloz_malakser(self):
        self.declare(PrzygotowanyMalakser(nazwa="tak"))
        print("Złożono malakser")

    @Rule(
        Malakser(),
        MalakserKomponent(nazwa="pokrywa"),
        MalakserKomponent(nazwa="ostrze dwustronne", strona=MATCH.strona),
        MalakserKomponent(nazwa="mocowanie na ostrze"),
        salience=2)
    def zamontuj_ostrze(self, strona):
        print("Zamontowano ostrze " + strona)


    #### Rozwiązania:
    @Rule(Rozwaizanie(nazwa=ROZWIAZANIE1))
    def rozwiazanie1(self):
        print("    Rozwiązanie: " + ROZWIAZANIE1)

    @Rule(Rozwaizanie(nazwa=ROZWIAZANIE2))
    def rozwiazanie2(self):
        print("    Rozwiązanie: " + ROZWIAZANIE2)

    @Rule(Rozwaizanie(nazwa=ROZWIAZANIE3))
    def rozwiazanie3(self):
        print("    Rozwiązanie: " + ROZWIAZANIE3)

    #### Przyczyny:
    @Rule(Przyczyna(nazwa=PRZYCZYNA1))
    def przyczyna1(self):
        print("  Potencjalna przyczyna: " + PRZYCZYNA1)
        self.declare(Rozwaizanie(nazwa=ROZWIAZANIE1))

    @Rule(Przyczyna(nazwa=PRZYCZYNA2))
    def przyczyna2(self):
        print("  Potencjalna przyczyna: " + PRZYCZYNA2)
        self.declare(Rozwaizanie(nazwa=ROZWIAZANIE2))

    @Rule(Przyczyna(nazwa=PRZYCZYNA3))
    def przyczyna3(self):
        print("  Potencjalna przyczyna: " + PRZYCZYNA3)
        self.declare(Rozwaizanie(nazwa=ROZWIAZANIE2))

    @Rule(Przyczyna(nazwa=PRZYCZYNA4))
    def przyczyna4(self):
        print("  Potencjalna przyczyna: " + PRZYCZYNA4)
        self.declare(Rozwaizanie(nazwa=ROZWIAZANIE3))

    #### PROBLEMY:
    @Rule(Problem(nazwa=PROBLEM1))
    def problem1(self):
        print("Wystąpił problem: " + PROBLEM1)
        self.declare(Przyczyna(nazwa=PRZYCZYNA1))
        self.declare(Przyczyna(nazwa=PRZYCZYNA2))

    @Rule(Problem(nazwa=PROBLEM2))
    def problem2(self):
        print("Wystąpił problem: " + PROBLEM2)
        self.declare(Przyczyna(nazwa=PRZYCZYNA3))
        self.declare(Przyczyna(nazwa=PRZYCZYNA4))

    @Rule(Problem(nazwa=PROBLEM3))
    def problem3(self):
        print("Wystąpił problem: " + PROBLEM3)
        self.declare(Przyczyna(nazwa=PRZYCZYNA3))
        self.declare(Przyczyna(nazwa=PRZYCZYNA4))

engine = MalakserExpert()

def przygotuj_salatke_z_uzyciem_popychacza():
    print("\nPrzygotuj sałatkę z użyciem popychacza:\n")
    engine.reset()

    # Zmontuj malakser
    engine.declare(Malakser())
    engine.declare(
    MalakserKomponent(nazwa="pokrywa"),
    MalakserKomponent(nazwa="mocowanie na ostrze"),
    MalakserKomponent(nazwa="pojemnik"),
    MalakserKomponent(nazwa="ostrze dwustronne", strona=STRONA_OSTRZA.TNACA),
    )
    # przygotuj sałatkę
    engine.declare(
        Skladnik(nazwa={"owoce", "warzywa"}),
        MalakserKomponent(nazwa="kabel zasilajacy", stan=ZASILANIE.JEST),
        MalakserKomponent(nazwa="włącznik", wlaczony=True),
        MalakserKomponent(nazwa="popychacz"),
    )
    engine.run()
    print("----------------------------------\n")

# Znajdz rozwiazanie problemu
def znajdz_rozwiazanie_problemu(problem):
    print("\nSzukam rozwiązań dla problemu: " + problem + ':\n')

    engine.reset()
    engine.declare(Problem(nazwa=problem))
    engine.run()
    print("----------------------------------")

# przygotuj_salatke_z_uzyciem_popychacza()

znajdz_rozwiazanie_problemu(PROBLEM1)
# znajdz_rozwiazanie_problemu(PROBLEM2)
# znajdz_rozwiazanie_problemu(PROBLEM3)























