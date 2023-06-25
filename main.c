#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <wchar.h>
#include <locale.h>

#define MAX_QUOTE_LENGTH 500
#define MAX_AUTHOR_LENGTH 100

int main() {
    setlocale(LC_ALL, ""); 

    FILE* file = fopen("./utils/quote.txt", "r");
    if (file == NULL) {
        printf("Error opening file.\n");
        return 1;
    }

    wchar_t quote[MAX_QUOTE_LENGTH];
    wchar_t author[MAX_AUTHOR_LENGTH];

    if (fgetws(quote, sizeof(quote) / sizeof(quote[0]), file) == NULL) {
        printf("Error reading quote from file.\n");
        fclose(file);
        return 1;
    }

    if (fgetws(author, sizeof(author) / sizeof(author[0]), file) == NULL) {
        printf("Error reading author from file.\n");
        fclose(file);
        return 1;
    }

    // Remove the newline character at the end of the strings
    quote[wcslen(quote) - 1] = L'\0';
    author[wcslen(author) - 1] = L'\0';

    wprintf(L"\nRandom Quote\n");
    wprintf(L"\n%ls\n\n%ls\n", quote, author);

    fclose(file);
    return 0;
}
