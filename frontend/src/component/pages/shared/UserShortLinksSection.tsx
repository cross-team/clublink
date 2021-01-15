import React, { Component } from 'react';
import { PageControl } from '../../ui/PageControl';
import { Table } from '../../ui/Table';
import { ShortLink } from '../../../entity/ShortLink';
import { IPagedShortLinks } from '../../../service/ShortLink.service';
import './UserShortLinksSection.scss';
import { Section } from '../../ui/Section';

interface IProps {
  pagedShortLinks?: IPagedShortLinks;
  pageSize: number;
  onPageLoad: (offset: number, pageSize: number) => void;
}

const DEFAULT_PROPS = {
  pageSize: 10
};

export class UserShortLinksSection extends Component<IProps> {
  static defaultProps: Partial<IProps> = DEFAULT_PROPS;

  componentDidMount(): void {
    this.showPage(0);
  }

  render() {
    if (!this.props.pagedShortLinks) {
      return false;
    }

    if (this.props.pagedShortLinks.totalCount < 1) {
      return false;
    }

    return (
      <div className={'UserShortLinksSection'}>
        <Section title={'Favorites'}>
          <Table
            headers={['Long Link', 'Alias']}
            rows={this.createTableRows()}
            widths={['70%', '30%']}
          />
          <div className={'page-control-wrapper'}>
            <PageControl
              totalPages={this.calculateTotalPages()}
              onPageChanged={this.onPageChanged}
            />
          </div>
        </Section>
      </div>
    );
  }

  private createTableRows = () => {
    const { shortLinks } = this.props.pagedShortLinks!;
    return shortLinks.map((shortLink: ShortLink) => {
      return [
        this.renderLongLink(shortLink.longLink),
        this.renderAlias(shortLink.alias)
      ];
    });
  };

  private renderLongLink = (longLink: string) => {
    return (
      <a
        className={'long-link'}
        href={longLink}
        target="_blank"
        rel="noopener noreferrer"
      >
        {longLink}
      </a>
    );
  };

  private renderAlias = (alias: string) => {
    return <span className={'alias'}>{alias}</span>;
  };

  public onPageChanged = (currentPageIdx: number) => {
    this.showPage(currentPageIdx);
  };

  private showPage = (pageIdx: number) => {
    const { pageSize } = this.props;
    const offset = pageIdx * pageSize;

    this.props.onPageLoad(offset, pageSize);
  };

  private calculateTotalPages = () => {
    const totalShortLinksCount = this.props.pagedShortLinks!.totalCount;
    return Math.ceil(totalShortLinksCount / this.props.pageSize);
  };
}
