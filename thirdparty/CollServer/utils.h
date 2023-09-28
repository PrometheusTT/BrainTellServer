﻿#ifndef UTILS_H
#define UTILS_H

#include <QCoreApplication>
#include <QDir>
#include "neuron_editing/neuron_format_converter.h"
#include <hiredis/hiredis.h>
const int neuron_type_color[21][3] = {
        {255, 255, 255},  // white,   0-undefined
        {20,  20,  20 },  // black,   1-soma
        {200, 20,  0  },  // red,     2-axon
        {0,   20,  200},  // blue,    3-dendrite
        {200, 0,   200},  // purple,  4-apical dendrite
        //the following is Hanchuan's extended color. 090331
        {0,   200, 200},  // cyan,    5
        {220, 200, 0  },  // yellow,  6
        {0,   200, 20 },  // green,   7
        {188, 94,  37 },  // coffee,  8
        {180, 200, 120},  // asparagus,	9
        {250, 100, 120},  // salmon,	10
        {120, 200, 200},  // ice,		11
        {100, 120, 200},  // orchid,	12
    //the following is Hanchuan's further extended color. 111003
    {255, 128, 168},  //	13
    {128, 255, 168},  //	14
    {128, 168, 255},  //	15
    {168, 255, 128},  //	16
    {255, 168, 128},  //	17
    {168, 128, 255}, //	18
    {0, 0, 0}, //19 //totally black. PHC, 2012-02-15
    //the following (20-275) is used for matlab heat map. 120209 by WYN
    {0,0,131}, //20
};

void dirCheck(QString dirBaseName);
QStringList getSwcInBlock(const QString msg,const V_NeuronSWC_list& testVNL);
QStringList getApoInBlock(const QString msg,const QList <CellAPO>& wholePoint);

void setredis(const int port,const char *ano);
void setexpire(const int port,const char *ano,const int expiretime);
void recoverPort(const int port);

vector<V_NeuronSWC>::iterator findseg(vector<V_NeuronSWC>::iterator begin,vector<V_NeuronSWC>::iterator end,const V_NeuronSWC seg);
NeuronTree convertMsg2NT(QStringList pointlist,int client,int user, int isMany, int mode=0);

double distance(const CellAPO &m1,const CellAPO &m2);
double distance(const double x1, const double x2, const double y1, const double y2, const double z1, const double z2);
double getSegLength(V_NeuronSWC &seg);
int findnearest(const CellAPO &m,const QList<CellAPO> &markers);

void init();
void stringToXYZ(string xyz, float& x, float& y, float& z);
#endif // UTILS_H