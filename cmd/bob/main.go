//All copyrights recived to MouhcineFD

#include <stdio.h>

#include <stdlib.h>

#include <time.h>

char tab[3][3],tabb[2][10];

int i,j,posij,nbrj,nbrp,pgan;

void afich();

void jouerhasar();

int scangan();

void utijouer();

int intljouer();

void rmtazero();

void pausesc(int nbresec)

{

    clock_t goal;

    goal=(nbresec*CLOCKS_PER_SEC)+clock();

    while(goal>clock())

        ;

}

int main()

{

    system("title X O By MouhcineFD");

    printf("////////////////////////////////////////////////////\n");

    printf("////////////////////////////////////////////////////\n");

    printf("//// Bienvenue Sur X O Devlopper Par MouhcineFD ////\n");

    printf("////////////////////////////////////////////////////\n");

    printf("////////////////////////////////////////////////////\n");

    pausesc(3);

    int choixm=1,choixm2,qint,choixm3,qjp;

    srand(time(NULL));

    while(choixm==1)

    {

        do

        {

            printf("\nCombien de partie voulez vous jouer ? ! Le maximan 10 !\n");

            scanf("%d",&nbrp);

            if(nbrp<=0||nbrp>10)

                printf("\nT'as choisi le mauvais choix !!!\nRepete la choix.\n");

        }

        while(nbrp<=0||nbrp>10);

        pgan=qjp=nbrp;

        for(i=0; i<3; i++)

        {

            for(j=0; j<3; j++)

            {

                tab[i][j]=' ';

            }

        }

        for(i=0; i<2; i++)

        {

            for(j=0; j<nbrp; j++)

            {

                tabb[i][j]=' ';

            }

        }

        nbrj=0;

        do

        {

            printf("\nVoulez-vous jouer avec ?\n1 : Ordinateur.\n2 : Un autre joueur.\n");

            scanf("%d",&choixm3);

            switch(choixm3)

            {

            case 1:

                do

                {

                    printf("\nSelectionnez le niveau :\n1 : Facile.\n2 : Difficile.\n");

                    scanf("%d",&choixm2);

                    switch(choixm2)

                    {

                    case 1:

                        while(nbrp>0)

                        {

                            system("cls");

                            if(qjp%2==0)

                            {

                                jouerhasar();

                                nbrj++;

                            }

                            while(nbrj<9&&scangan()!=0)

                            {

                                if(nbrj<9&&scangan()!=0)

                                {

                                    printf("\n Vous (O).\n");

                                    utijouer();

                                    nbrj++;

                                }

                                if(nbrj<9&&scangan()!=0)

                                {

                                    if(casajouer()!=0)

                                        jouerhasar();

                                    nbrj++;

                                }

                            }

                            rmtazero();

                            nbrp--;

                            qjp=nbrp;

                        }

                        break;

                    case 2:



                        while(nbrp>0)

                        {

                            system("cls");

                            tab[1][1]='X';

                            utijouer();

                            if(intljouer()!=0)

                                if(casajouer()!=0)

                                    jouerhasar();

                            nbrj=5;

                            while(nbrj<9&&scangan()!=0)

                            {

                                utijouer();

                                nbrj++;

                                if(nbrj<9&&scangan()!=0)

                                {

                                    if(casajouer()!=0)

                                        jouerhasar();

                                    nbrj++;

                                }

                            }

                            rmtazero();

                            nbrp--;

                        }

                        break;

                    default:

                        printf("\nT'as choisi le mauvais choix !!!\nRepete la choix.\n");

                        break;

                    }

                }

                while(choixm2!=1&&choixm2!=2);

            case 2:

                while(nbrp>0)

                {

                    system("cls");

                    while(nbrj<9&&scangan()!=0)

                    {

                        if(qjp%2==0)

                        {

                            if(nbrj<9&&scangan()!=0)

                            {

                                printf("\nLe role de joueur "" O "" .\n");

                                utijouer();

                                nbrj++;

                            }

                        }

                        if(nbrj<9&&scangan()!=0)

                        {

                            printf("\nLe role de joueur "" X "" .\n");

                            utijouer();

                            tab[(posij%10)-1][(posij/10)-1]='X';

                            nbrj++;

                        }

                        qjp=2;

                    }

                    rmtazero();

                    nbrp--;

                    qjp=nbrp;

                }

                break;

            default:

                printf("\nT'as choisi le mauvais choix !!!\nRepete la choix.\n");

                break;

            }

        }

        while(choixm3!=1&&choixm3!=2);

        system("cls");

        for(i=0; i<3; i++)

        {

            for(j=0; j<3; j++)

            {

                tab[i][j]=' ';

            }

        }

        afich();

        choixm2=choixm3=0;

        for(i=0; i<2; i++)

        {

            for(j=0; j<pgan; j++)

            {

                if(tabb[i][j]=='1')

                {

                    if(i==0)

                        choixm2++;

                    if(i==1)

                        choixm3++;



                }

            }

        }

        if(choixm2>choixm3)

        {

            printf("\n\n***********************************************\n");

            printf("*** Le gagnant au total parties est le ** X ***\n");

            printf("***********************************************\n\n");

        }

        else if(choixm2<choixm3)

        {

            printf("\n\n***********************************************\n");

            printf("*** Le gagnant au total parties est le ** O ***\n");

            printf("***********************************************\n\n");

        }

        else

        {

            printf("\n\n*****************************************************\n");

            printf("*** L'egalite entre les joueurs au total parties. ***\n");

            printf("*****************************************************\n\n");

        }

        system("pause");

        system("cls");

        do

        {

            printf("\nVoulez-vous jouer a nouveau ?\n1 : OUI\n2 : NON\n");

            scanf("%d",&choixm);

            if(choixm!=1&&choixm!=2)

                printf("\nT'as choisi le mauvais choix !!!\nRepete la choix.\n");

        }

        while(choixm!=1&&choixm!=2);

        system("cls");

    }

    printf("////////////////////////////////////////\n");

    printf("////////////// MouhcineFD //////////////\n");

    printf("//// vous remercions d'avoir choisi ////\n");

    printf("///////////////// X O //////////////////\n");

    printf("////////////////////////////////////////\n");

    pausesc(3);

    return 0;

}

void rmtazero()

{

    afich();

    if(scangan()==0)

    {

        printf("\nLe gagnant dans cette partie est le "" %c "" .\n\n",tab[i][j]);

        if(tab[i][j]=='X')

        {

            tabb[0][pgan-nbrp]='1';

            tabb[1][pgan-nbrp]='0';

        }

        else

        {

            tabb[1][pgan-nbrp]='1';

            tabb[0][pgan-nbrp]='0';

        }

    }

    else

    {

        printf("\nL'egalite entre les joueurs au cette partie.\n\n");

        tabb[1][pgan-nbrp]='0';

        tabb[0][pgan-nbrp]='0';

    }

    system("pause");

    for(i=0; i<3; i++)

    {

        for(j=0; j<3; j++)

        {

            tab[i][j]=' ';

        }

    }

    nbrj=0;

}

void afich()

{

    int v;

    printf("\n     X1  X2  X3\n");

    printf("    ___________              ");

    for(v=0; v<pgan+1; v++)

    {

        printf("____");

    }

    printf("\n");

    for(i=0; i<3; i++)

    {

        printf("Y%d ",i+1);

        for(j=0; j<3; j++)

        {

            printf("| %c ",tab[i][j]);

        }



        printf("|");

        printf("             |");

        if(i==0)

        {

            for(v=0; v<pgan+1; v++)

            {

                if(v==0)

                    printf("   |");

                else

                {

                    if(v<10)

                        printf("P%d |",v);

                    else

                        printf("P%d|",v);

                }

            }

        }

        else

        {

            for(v=0; v<pgan+1; v++)

            {

                if(v==0&&i==1)

                    printf(" X |");

                else if(v==0&&i==2)

                    printf(" O |");

                else

                    printf(" %c |",tabb[i-1][v-1]);

            }

        }

        printf("\n   |___|___|___|             |");

        for(v=0; v<pgan+1; v++)

        {

            printf("___|");

        }

        printf("\n");

    }

}

void jouerhasar()

{

    do

    {

        i=(rand()%(2-0+1))+0;

        j=(rand()%(2-0+1))+0;

    }

    while(tab[i][j]!=' ');

    tab[i][j]='X';

}

int scangan()

{

    for(i=j=0; j<3; j++)

    {

        if(tab[i][j]!=' ')

            if(tab[i][j]==tab[i+1][j]&&tab[i+1][j]==tab[i+2][j])



                return 0;



    }

    for(i=j=0; i<3; i++)

    {

        if(tab[i][j]!=' ')

            if(tab[i][j]==tab[i][j+1]&&tab[i][j+1]==tab[i][j+2])

                return 0;

    }

    if(tab[j+1][j+1]!=' ')

    {

        if(tab[j][j]==tab[j+1][j+1]&&tab[j+1][j+1]==tab[j+2][j+2])

        {

            tab[i][j]=tab[j][j];

            return 0;

        }

        if(tab[j+2][j]==tab[j+1][j+1]&&tab[j+1][j+1]==tab[j][j+2])

        {

            tab[i][j]=tab[j+2][j];

            return 0;

        }

    }

}

int casajouer()

{

    for(i=j=0; j<3; j++)

    {

        if(tab[i][j]==' ')

            if(tab[i+1][j]=='X'&&tab[i+1][j]==tab[i+2][j])

            {

                tab[i][j]='X';

                return 0;

            }

        if(tab[i+1][j]==' ')

            if(tab[i][j]=='X'&&tab[i][j]==tab[i+2][j])

            {

                tab[i+1][j]='X';

                return 0;

            }

        if(tab[i+2][j]==' ')

            if(tab[i][j]=='X'&&tab[i][j]==tab[i+1][j])

            {

                tab[i+2][j]='X';

                return 0;

            }

    }

    for(i=j=0; i<3; i++)

    {

        if(tab[i][j]==' ')

            if(tab[i][j+1]=='X'&&tab[i][j+1]==tab[i][j+2])

            {

                tab[i][j]='X';

                return 0;

            }

        if(tab[i][j+1]==' ')

            if(tab[i][j]=='X'&&tab[i][j]==tab[i][j+2])

            {

                tab[i][j+1]='X';

                return 0;

            }

        if(tab[i][j+2]==' ')

            if(tab[i][j]=='X'&&tab[i][j]==tab[i][j+1])

            {

                tab[i][j+2]='X';

                return 0;

            }

    }

    i=j=0;

    if(tab[i][j]==' ')

        if(tab[i+1][j+1]=='X'&&tab[i+1][j+1]==tab[i+2][j+2])

        {

            tab[i][j]='X';

            return 0;

        }

    if(tab[i+1][j+1]==' ')

        if(tab[i][j]=='X'&&tab[i][j]==tab[i+2][j+2])

        {

            tab[i+1][j+1]='X';

            return 0;

        }

    if(tab[i+2][j+2]==' ')

        if(tab[i][j]=='X'&&tab[i][j]==tab[i+1][j+1])

        {

            tab[i+2][j+2]='X';

            return 0;

        }

    if(tab[i][j+2]==' ')

        if(tab[i+1][j+1]=='X'&&tab[i+1][j+1]==tab[i+2][j])

        {

            tab[i][j+2]='X';

            return 0;

        }

    if(tab[i+1][j+1]==' ')

        if(tab[i][j+2]=='X'&&tab[i][j]==tab[i+2][j])

        {

            tab[i+1][j+1]='X';

            return 0;

        }

    if(tab[i+2][j]==' ')

        if(tab[i][j+2]=='X'&&tab[i][j+2]==tab[i+1][j+1])

        {

            tab[i+2][j]='X';

            return 0;

        }

    for(i=j=0; j<3; j++)

    {

        if(tab[i][j]==' ')

            if(tab[i+1][j]=='O'&&tab[i+1][j]==tab[i+2][j])

            {

                tab[i][j]='X';

                return 0;

            }

        if(tab[i+1][j]==' ')

            if(tab[i][j]=='O'&&tab[i][j]==tab[i+2][j])

            {

                tab[i+1][j]='X';

                return 0;

            }

        if(tab[i+2][j]==' ')

            if(tab[i][j]=='O'&&tab[i][j]==tab[i+1][j])

            {

                tab[i+2][j]='X';

                return 0;

            }

    }

    for(i=j=0; i<3; i++)

    {

        if(tab[i][j]==' ')

            if(tab[i][j+1]=='O'&&tab[i][j+1]==tab[i][j+2])

            {

                tab[i][j]='X';

                return 0;

            }

        if(tab[i][j+1]==' ')

            if(tab[i][j]=='O'&&tab[i][j]==tab[i][j+2])

            {

                tab[i][j+1]='X';

                return 0;

            }

        if(tab[i][j+2]==' ')

            if(tab[i][j]=='O'&&tab[i][j]==tab[i][j+1])

            {

                tab[i][j+2]='X';

                return 0;

            }

    }

    i=j=0;

    if(tab[i][j]==' ')

        if(tab[i+1][j+1]=='O'&&tab[i+1][j+1]==tab[i+2][j+2])

        {

            tab[i][j]='X';

            return 0;

        }

    if(tab[i+1][j+1]==' ')

        if(tab[i][j]=='O'&&tab[i][j]==tab[i+2][j+2])

        {

            tab[i+1][j+1]='X';

            return 0;

        }

    if(tab[i+2][j+2]==' ')

        if(tab[i][j]=='O'&&tab[i][j]==tab[i+1][j+1])

        {

            tab[i+2][j+2]='X';

            return 0;

        }

    if(tab[i][j+2]==' ')

        if(tab[i+1][j+1]=='O'&&tab[i+1][j+1]==tab[i+2][j])

        {

            tab[i][j+2]='X';

            return 0;

        }

    if(tab[i+1][j+1]==' ')

        if(tab[i][j+2]=='O'&&tab[i][j]==tab[i+2][j])

        {

            tab[i+1][j+1]='X';

            return 0;

        }

    if(tab[i+2][j]==' ')

        if(tab[i][j+2]=='O'&&tab[i][j+2]==tab[i+1][j+1])

        {

            tab[i+2][j]='X';

            return 0;

        }

}

void utijouer()

{

    afich();

    do

    {

        printf("\nEntrez les coordonnees de la case a jouer sur la forme XY : ");

        scanf("%d",&posij);

        if(tab[(posij%10)-1][(posij/10)-1]==' ')

            tab[(posij%10)-1][(posij/10)-1]='O';

        else

        {

            printf("\nT'as choisi le mauvais case !!!\nRepete la choix.\n");

            posij=0;

        }

    }

    while(posij==0);

    system("cls");

}

int intljouer()

{

    if(tab[0][0]=='O')

    {

        tab[2][2]='X';

        utijouer();

        if(tab[2][1]=='O')

        {

            tab[1][2]='X';

            return 0;

        }

        else if(tab[1][2]=='O')

        {

            tab[2][1]='X';

            return 0;

        }

    }

    else if(tab[2][2]=='O')

    {

        tab[0][0]='X';

        utijouer();

        if(tab[1][0]=='O')

        {

            tab[0][1]='X';

            return 0;

        }

        else if(tab[0][1]=='O')

        {

            tab[1][0]='X';

            return 0;

        }

    }

    else if(tab[2][0]=='O')

    {

        tab[0][2]='X';

        utijouer();

        if(tab[1][2]=='O')

        {

            tab[0][1]='X';

            return 0;

        }

        else if(tab[0][1]=='O')

        {

            tab[1][2]='X';

            return 0;

        }

    }

    else if(tab[0][2]=='O')

    {

        tab[2][0]='X';

        utijouer();

        if(tab[1][0]=='O')

        {

            tab[2][1]='X';

            return 0;

        }

        else if(tab[2][1]=='O')

        {

            tab[1][0]='X';

            return 0;

        }

    }

    else if(tab[1][0]=='O')

    {

        tab[1][2]='X';

        utijouer();



        if(tab[2][1]=='O')

        {

            tab[2][2]='X';

            return 0;

        }

        else if(tab[0][1]=='O')

        {

            tab[0][2]='X';

            return 0;

        }

    }

    else if(tab[1][2]=='O')

    {

        tab[1][0]='X';

        utijouer();

        if(tab[0][1]=='O')

        {

            tab[0][0]='X';

            return 0;

        }

        else if(tab[2][1]=='O')

        {

            tab[2][0]='X';

            return 0;

        }

    }

    else if(tab[2][1]=='O')

    {

        tab[0][1]='X';

        utijouer();

        if(tab[1][0]=='O')

        {

            tab[0][0]='X';

            return 0;

        }

        else if(tab[1][2]=='O')

        {

            tab[0][2]='X';

            return 0;

        }

    }

    else if(tab[0][1]=='O')

    {

        tab[2][1]='X';

        utijouer();

        if(tab[1][0]=='O')

        {

            tab[2][0]='X';

            return 0;

        }

        else if(tab[1][2]=='O')

        {

            tab[2][2]='X';

            return 0;

        }

    }

}
